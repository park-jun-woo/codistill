//ff:type feature=scan type=test topic=quarkus
//ff:what classPathPrefixSynthesisSource 테스트 보조 선언(Phase194)
package quarkus

const classPathPrefixSynthesisSource = `
@Path("/api")
public class ItemResource {
    @GET
    @Path("/items")
    public ItemDto listItems() { return null; }

    @GET
    public ItemDto root() { return null; }
}
`
