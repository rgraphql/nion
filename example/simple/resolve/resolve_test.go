package resolve

import (
	"context"
	"testing"
	"time"

	"github.com/rgraphql/nion/example/simple"

	"github.com/rgraphql/nion/encoder"
	qttestutil "github.com/rgraphql/nion/qtree/testutil"
	"github.com/rgraphql/nion/resolver"
	proto "github.com/rgraphql/rgraphql"
)

// .\nion.exe analyze --schema ..\..\example\simple\schema.graphql --go-pkg "github.com/rgraphql/nion/example/simple" --go-query-type RootResolver --go-output "../../example/simple/resolve/resolve_generated.go"
func TestResolveSimple(t *testing.T) {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer ctxCancel()

	schema, qtNode, errCh := qttestutil.BuildMockTree(t)
	_ = schema

	qtNode.ApplyTreeMutation(&proto.RGQLQueryTreeMutation{
		NodeMutation: []*proto.RGQLQueryTreeMutation_NodeMutation{
			&proto.RGQLQueryTreeMutation_NodeMutation{
				NodeId:    0,
				Operation: proto.RGQLQueryTreeMutation_SUBTREE_ADD_CHILD,
				Node: &proto.RGQLQueryTreeNode{
					Id:        1,
					FieldName: "names",
				},
			},
			&proto.RGQLQueryTreeMutation_NodeMutation{
				NodeId:    0,
				Operation: proto.RGQLQueryTreeMutation_SUBTREE_ADD_CHILD,
				Node: &proto.RGQLQueryTreeNode{
					Id:        2,
					FieldName: "allPeople",
					Children: []*proto.RGQLQueryTreeNode{
						&proto.RGQLQueryTreeNode{
							Id:        3,
							FieldName: "name",
						},
						&proto.RGQLQueryTreeNode{
							Id:        4,
							FieldName: "height",
						},
					},
				},
			},
			&proto.RGQLQueryTreeMutation_NodeMutation{
				NodeId:    0,
				Operation: proto.RGQLQueryTreeMutation_SUBTREE_ADD_CHILD,
				Node: &proto.RGQLQueryTreeNode{
					Id:        5,
					FieldName: "singlePerson",
					Children: []*proto.RGQLQueryTreeNode{
						&proto.RGQLQueryTreeNode{
							Id:        6,
							FieldName: "name",
						},
					},
				},
			},
		},
	})

	go func() {
		err := <-errCh
		t.Fatal(err.GetError())
	}()

	encoder := encoder.NewResultEncoder(50)
	outputCh := make(chan []byte)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dat := <-outputCh:
				_ = dat
			}
		}
	}()
	go encoder.Run(ctx, outputCh)

	resolverCtx := resolver.NewContext(ctx, qtNode, encoder)
	rootRes := &simple.RootResolver{}
	ResolveRootQuery(resolverCtx, rootRes)
}
