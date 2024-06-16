// @generated by protobuf-ts 2.9.4 with parameter generate_dependencies
// @generated from protobuf file "grpc/health/v1/health.proto" (package "grpc.health.v1", syntax proto3)
// tslint:disable
//
// Copyright 2015 The gRPC Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//
// The canonical version of this proto can be found at
// https://github.com/grpc/grpc-proto/blob/master/grpc/health/v1/health.proto
//
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message grpc.health.v1.HealthCheckRequest
 */
export interface HealthCheckRequest {
    /**
     * @generated from protobuf field: string service = 1;
     */
    service: string;
}
/**
 * @generated from protobuf message grpc.health.v1.HealthCheckResponse
 */
export interface HealthCheckResponse {
    /**
     * @generated from protobuf field: grpc.health.v1.HealthCheckResponse.ServingStatus status = 1;
     */
    status: HealthCheckResponse_ServingStatus;
}
/**
 * @generated from protobuf enum grpc.health.v1.HealthCheckResponse.ServingStatus
 */
export enum HealthCheckResponse_ServingStatus {
    /**
     * @generated from protobuf enum value: UNKNOWN = 0;
     */
    UNKNOWN = 0,
    /**
     * @generated from protobuf enum value: SERVING = 1;
     */
    SERVING = 1,
    /**
     * @generated from protobuf enum value: NOT_SERVING = 2;
     */
    NOT_SERVING = 2,
    /**
     * Used only by the Watch method.
     *
     * @generated from protobuf enum value: SERVICE_UNKNOWN = 3;
     */
    SERVICE_UNKNOWN = 3
}
// @generated message type with reflection information, may provide speed optimized methods
class HealthCheckRequest$Type extends MessageType<HealthCheckRequest> {
    constructor() {
        super("grpc.health.v1.HealthCheckRequest", [
            { no: 1, name: "service", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<HealthCheckRequest>): HealthCheckRequest {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.service = "";
        if (value !== undefined)
            reflectionMergePartial<HealthCheckRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: HealthCheckRequest): HealthCheckRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string service */ 1:
                    message.service = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: HealthCheckRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string service = 1; */
        if (message.service !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.service);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message grpc.health.v1.HealthCheckRequest
 */
export const HealthCheckRequest = new HealthCheckRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class HealthCheckResponse$Type extends MessageType<HealthCheckResponse> {
    constructor() {
        super("grpc.health.v1.HealthCheckResponse", [
            { no: 1, name: "status", kind: "enum", T: () => ["grpc.health.v1.HealthCheckResponse.ServingStatus", HealthCheckResponse_ServingStatus] }
        ]);
    }
    create(value?: PartialMessage<HealthCheckResponse>): HealthCheckResponse {
        const message = globalThis.Object.create((this.messagePrototype!));
        message.status = 0;
        if (value !== undefined)
            reflectionMergePartial<HealthCheckResponse>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: HealthCheckResponse): HealthCheckResponse {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* grpc.health.v1.HealthCheckResponse.ServingStatus status */ 1:
                    message.status = reader.int32();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: HealthCheckResponse, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* grpc.health.v1.HealthCheckResponse.ServingStatus status = 1; */
        if (message.status !== 0)
            writer.tag(1, WireType.Varint).int32(message.status);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message grpc.health.v1.HealthCheckResponse
 */
export const HealthCheckResponse = new HealthCheckResponse$Type();
/**
 * @generated ServiceType for protobuf service grpc.health.v1.Health
 */
export const Health = new ServiceType("grpc.health.v1.Health", [
    { name: "Check", options: {}, I: HealthCheckRequest, O: HealthCheckResponse },
    { name: "Watch", serverStreaming: true, options: {}, I: HealthCheckRequest, O: HealthCheckResponse }
]);