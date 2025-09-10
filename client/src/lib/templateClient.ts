import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { TemplateServicesClient } from "../proto/Template.client";


const transport = new GrpcWebFetchTransport({
    baseUrl: '/rpc/s/'
})

const client = new TemplateServicesClient(transport);

export default client