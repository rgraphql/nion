package qtree

import (
	"errors"
	"testing"

	. "github.com/rgraphql/nion/qtree"
	"github.com/rgraphql/nion/qtree/testutil"
	"github.com/rgraphql/nion/schema"
	proto "github.com/rgraphql/rgraphql"
)

func buildMockTree(t *testing.T) (*schema.Schema, *QueryTreeNode, <-chan *proto.RGQLQueryError) {
	return testutil.BuildMockTree(t)
}

func TestBasics(t *testing.T) {
	_, qt, _ := buildMockTree(t)
	err := qt.AddChild(&proto.RGQLQueryTreeNode{
		Id:        1,
		FieldName: "allPeople",
		Children: []*proto.RGQLQueryTreeNode{
			{
				Id:        2,
				FieldName: "name",
			},
		},
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("%#v", qt.Children[0])
}

func TestSchemaErrors(t *testing.T) {
	_, qt, errCh := buildMockTree(t)
	qt.AddChild(&proto.RGQLQueryTreeNode{
		Id:        1,
		FieldName: "allPeople",
		Children: []*proto.RGQLQueryTreeNode{
			{
				Id:        2,
				FieldName: "names",
			},
		},
	})
	var err error
	select {
	case e := <-errCh:
		err = errors.New(e.Error)
	default:
	}
	if err == nil || err.Error() != "Invalid field names on Person." {
		t.Fatalf("Did not return expected error (%v).", err)
	}
}
