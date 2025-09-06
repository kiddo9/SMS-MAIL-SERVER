import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { TemplateServicesClient } from "../proto/Template.client";


const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENVOY_URL
})

const client = new TemplateServicesClient(transport);

export default client