package jsonapi

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestControllerCreation(t *testing.T) {
	cNew := New()

	assertNotEmpty(t, cNew.Models)
	assertEqual(t, "/", cNew.APIURLBase)
	assertEqual(t, 1, cNew.IncludeNestedLimit)

	cDefault := Default()
	assertNotEmpty(t, cDefault.Models)
	assertTrue(t, cDefault.APIURLBase == "/api")
	assertNotEqual(t, cNew.ErrorLimitMany, cDefault.ErrorLimitMany)
	assertNotEqual(t, cNew.ErrorLimitSingle, cDefault.ErrorLimitSingle)
}

func TestBuildScopeMany(t *testing.T) {
	var (
		err   error
		req   *http.Request
		scope *Scope
		errs  []*ErrorObject
		c     *Controller
	)

	c = Default()
	err = c.PrecomputeModels(&Blog{}, &Post{}, &Comment{})
	assertNil(t, err)

	// raw scope without query
	req = httptest.NewRequest("GET", "/api/v1/blogs", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertEmpty(t, errs)
	assertNil(t, err)
	assertNotNil(t, scope)

	assertNil(t, scope.Root)
	assertEmpty(t, scope.SubScopes)

	assertNotEmpty(t, scope.Fields)
	assertEmpty(t, scope.Filters)
	assertEmpty(t, scope.Sorts)
	assertNil(t, scope.PaginationScope)
	assertEqual(t, scope.Struct, c.MustGetModelStruct(&Blog{}))
	assertNil(t, scope.RelatedField)

	// with include
	req = httptest.NewRequest("GET", "/api/v1/blogs?include=current-post", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertEmpty(t, errs)
	assertNil(t, err)
	assertNotNil(t, scope)

	assertNotEmpty(t, scope.SubScopes)
	assertEqual(t, scope.SubScopes[0].Struct, c.MustGetModelStruct(&Post{}))

	c.PrecomputeModels(&Blog{}, &Post{}, &Comment{})

	// with sorts
	req = httptest.NewRequest("GET", "/api/v1/blogs?sort=id,-title,posts.id", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)

	assertNotNil(t, scope)

	assertEqual(t, 3, len(scope.Sorts))
	assertEqual(t, AscendingOrder, scope.Sorts[0].Order)
	assertEqual(t, "id", scope.Sorts[0].Field.jsonAPIName)
	assertEqual(t, DescendingOrder, scope.Sorts[1].Order)
	assertEqual(t, "title", scope.Sorts[1].Field.jsonAPIName)
	assertEqual(t, "posts", scope.Sorts[2].Field.jsonAPIName)
	assertEqual(t, 1, len(scope.Sorts[2].RelScopes))
	assertEqual(t, AscendingOrder, scope.Sorts[2].RelScopes[0].Order)
	assertEqual(t, "id", scope.Sorts[2].RelScopes[0].Field.jsonAPIName)

	req = httptest.NewRequest("GET", "/api/v1/blogs?sort=posts.id,posts.title", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)

	assertEqual(t, 1, len(scope.Sorts))
	assertEqual(t, 2, len(scope.Sorts[0].RelScopes))
	assertEqual(t, "id", scope.Sorts[0].RelScopes[0].Field.jsonAPIName)
	assertEqual(t, "title", scope.Sorts[0].RelScopes[1].Field.jsonAPIName)

	// paginations
	req = httptest.NewRequest("GET", "/api/v1/blogs?page[size]=4&page[number]=5", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)
	assertNotNil(t, scope)

	assertNotNil(t, scope.PaginationScope)
	assertEqual(t, 4, scope.PaginationScope.PageSize)
	assertEqual(t, 5, scope.PaginationScope.PageNumber)

	// pagination limit, offset
	req = httptest.NewRequest("GET", "/api/v1/blogs?page[limit]=10&page[offset]=5", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)
	assertNotNil(t, scope)

	assertNotNil(t, scope.PaginationScope)
	assertEqual(t, 10, scope.PaginationScope.Limit)
	assertEqual(t, 5, scope.PaginationScope.Offset)

	// pagination errors
	req = httptest.NewRequest("GET", "/api/v1/blogs?page[limit]=have&page[offset]=a&page[size]=nice&page[number]=day", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)
	// t.Log(errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs?page[limit]=2&page[number]=1", nil)
	_, errs, _ = c.BuildScopeMany(req, &Blog{})
	assertNotEmpty(t, errs)

	// filter
	req = httptest.NewRequest("GET", "/api/v1/blogs?filter[blogs][id][eq]=12,55", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)
	assertNotNil(t, scope)

	assertEqual(t, 1, len(scope.Filters))
	assertEqual(t, OpEqual, scope.Filters[0].Values[0].Operator)
	assertEqual(t, 2, len(scope.Filters[0].Values[0].Values))

	// invalid filter
	//	- invalid bracket
	//	- invalid operator
	//	- invalid value
	//	- not included collection - 'posts'
	req = httptest.NewRequest("GET", "/api/v1/blogs?filter[[blogs][id][eq]=12,55&filter[blogs][id][invalid]=125&filter[blogs][id]=stringval&filter[posts][id]=12&fields[blogs]=id", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs?filter[blogs]=somethingnotid&filter[blogs][id]=againbad&filter[blogs][posts][id]=badid", nil)
	_, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	c.PrecomputeModels(&Blog{}, &Post{}, &Comment{})

	// fields
	req = httptest.NewRequest("GET", "/api/v1/blogs?fields[blogs]=title,posts", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)

	// title, posts and id
	assertEqual(t, 3, len(scope.Fields))
	assertNotEqual(t, scope.Fields[0].fieldName, scope.Fields[1].fieldName)

	// fields error
	//	- bracket error
	//	- nested error
	//	- invalid collection name
	req = httptest.NewRequest("GET", "/api/v1/blogs?fields[[blogs]=title&fields[blogs][title]=now&fields[blog]=title&fields[blogs]=title&fields[blogs]=posts", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	// field error too many
	req = httptest.NewRequest("GET", "/api/v1/blogs?fields[blogs]=title,id,posts,comments,this-comment,some-invalid,current-post", nil)
	_, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	// sorterror
	req = httptest.NewRequest("GET", "/api/v1/blogs?sort=posts.comments.id,current-post.itle,postes.comm", nil)
	_, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	// unsupported parameter
	req = httptest.NewRequest("GET", "/api/v1/blogs?title=name", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	// too many errors
	// after 5 errors the function stops
	req = httptest.NewRequest("GET", "/api/v1/blogs?fields[[blogs]=title&fields[blogs][title]=now&fields[blog]=title&sort=-itle&filter[blog][id]=1&filter[blogs][unknown]=123&filter[blogs][current-post][something]=123", nil)
	scope, errs, err = c.BuildScopeMany(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	//internal
	req = httptest.NewRequest("GET", "/api/v1/blogs", nil)
	c.Models.Set(reflect.TypeOf(Blog{}), nil)
	_, _, err = c.BuildScopeMany(req, &Blog{})
	assertError(t, err)
}

func TestBuildScopeSingle(t *testing.T) {
	c := Default()
	err := c.PrecomputeModels(&Blog{}, &Post{}, &Comment{})
	assertNil(t, err)

	req := httptest.NewRequest("GET", "/api/v1/blogs/55", nil)
	scope, errs, err := c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)
	assertNotNil(t, scope)

	assertEqual(t, 55, scope.Filters[0].Values[0].Values[0])

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[posts]=title", nil)
	scope, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertEmpty(t, errs)
	assertNotNil(t, scope)

	assertEqual(t, 44, scope.Filters[0].Values[0].Values[0])
	assertEqual(t, 1, len(scope.SubScopes))
	assertEqual(t, 1, len(scope.SubScopes[0].Fields))

	// errored
	req = httptest.NewRequest("GET", "/api/v1/blogs", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertError(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/posts/1", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertError(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/bad-id", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=invalid", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[blogs]=title,posts&fields[blogs]=posts", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[blogs]]=posts", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[blogs]=title,posts&fields[blogs][posts]=title", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[blogs]=title,posts&fields[blogs]=posts", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/44?include=posts&fields[postis]=title", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/123?title=some-title", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)

	req = httptest.NewRequest("GET", "/api/v1/blogs/123?fields[postis]=title&fields[posts]=idss&fields[posts]=titles&title=sometitle&fields[blogs]=titles,current-posts", nil)
	_, errs, err = c.BuildScopeSingle(req, &Blog{})
	assertNil(t, err)
	assertNotEmpty(t, errs)
	t.Log(len(errs))
	t.Log(errs)
}

func TestPrecomputeModels(t *testing.T) {
	// input valid models
	validModels := []interface{}{&User{}, &Pet{}, &Driver{}, &Car{}, &WithPointer{}, &Blog{}, &Post{}, &Comment{}}
	err := c.PrecomputeModels(validModels...)
	if err != nil {
		t.Error(err)
	}
	clearMap()

	// if somehow map is nil
	c.Models = nil
	err = c.PrecomputeModels(&Timestamp{})
	if err != nil {
		t.Error(err)
	}

	// if one of the relationship is not precomputed
	clearMap()
	// User has relationship with Pet
	err = c.PrecomputeModels(&User{})
	if err == nil {
		t.Error("The User is related to Pets and so that should be an error")
	}
	clearMap()

	// if one of the models is invalid
	err = c.PrecomputeModels(&Timestamp{}, &BadModel{})
	if err == nil {
		t.Error("BadModel should not be accepted in precomputation.")
	}

	// provided Struct type to precompute models
	err = c.PrecomputeModels(Timestamp{})
	if err == nil {
		t.Error("A pointer to the model should be provided.")
	}

	// provided ptr to basic type
	basic := "value"
	err = c.PrecomputeModels(&basic)
	if err == nil {
		t.Error("Only structs should be accepted!")
	}

	// provided slice
	err = c.PrecomputeModels(&[]*Timestamp{})
	if err == nil {
		t.Error("Slice should not be accepted in precomputedModels")
	}

	// if no tagged fields are provided an error would be thrown
	err = c.PrecomputeModels(&ModelNonTagged{})
	if err == nil {
		t.Error("Non tagged models are not allowed.")
	}
	clearMap()

	// models without primary are not allowed.
	err = c.PrecomputeModels(&NoPrimaryModel{})
	if err == nil {
		t.Error("No primary field provided.")
	}
	clearMap()

	type InvalidPrimaryField struct {
		ID float64 `jsonapi:"primary,invalids"`
	}

	err = c.PrecomputeModels(&InvalidPrimaryField{})
	assertError(t, err)
}

func TestGetModelStruct(t *testing.T) {
	// MustGetModelStruct
	// if the model is not in the cache map
	clearMap()
	assertPanic(t, func() {
		c.MustGetModelStruct(Timestamp{})
	})

	c.Models.Set(reflect.TypeOf(Timestamp{}), &ModelStruct{})
	mStruct := c.MustGetModelStruct(Timestamp{})
	if mStruct == nil {
		t.Error("The model struct shoud not be nil.")
	}

	// GetModelStruct
	// providing ptr should return mStruct
	var err error
	_, err = c.GetModelStruct(&Timestamp{})
	if err != nil {
		t.Error(err)
	}

	// nil model
	_, err = c.GetModelStruct(nil)
	if err == nil {
		t.Error(err)
	}
}
