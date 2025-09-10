import { GrpcWebFetchTransport } from "@protobuf-ts/grpcweb-transport";
import { fileUploadServicesClient } from "../proto/FileUpload.client";


const transport = new GrpcWebFetchTransport({
    baseUrl: import.meta.env.VITE_ENV === "development" ? import.meta.env.VITE_ENVOY_URL : '/rpc/s/'
})

const client = new fileUploadServicesClient(transport);

export default client