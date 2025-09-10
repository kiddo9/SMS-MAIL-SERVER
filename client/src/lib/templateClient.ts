import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { TemplateServicesClient } from "../proto/Template.client";


const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENV === "development" ? import.meta.env.VITE_ENVOY_URL : '/rpc/s/'
})

const client = new TemplateServicesClient(transport);

export default client