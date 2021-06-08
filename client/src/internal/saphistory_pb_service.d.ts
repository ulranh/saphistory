// package: private
// file: internal/saphistory.proto

import * as internal_saphistory_pb from "../internal/saphistory_pb";
import {grpc} from "@improbable-eng/grpc-web";

type SapHistoryServiceGetSapStatus = {
  readonly methodName: string;
  readonly service: typeof SapHistoryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof internal_saphistory_pb.SapSelection;
  readonly responseType: typeof internal_saphistory_pb.TransactionData;
};

type SapHistoryServiceGetSystemList = {
  readonly methodName: string;
  readonly service: typeof SapHistoryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof internal_saphistory_pb.Nothing;
  readonly responseType: typeof internal_saphistory_pb.SystemList;
};

type SapHistoryServiceUpdateSystem = {
  readonly methodName: string;
  readonly service: typeof SapHistoryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof internal_saphistory_pb.SystemInfo;
  readonly responseType: typeof internal_saphistory_pb.Nothing;
};

type SapHistoryServiceDeleteSystem = {
  readonly methodName: string;
  readonly service: typeof SapHistoryService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof internal_saphistory_pb.SystemInfo;
  readonly responseType: typeof internal_saphistory_pb.Nothing;
};

export class SapHistoryService {
  static readonly serviceName: string;
  static readonly GetSapStatus: SapHistoryServiceGetSapStatus;
  static readonly GetSystemList: SapHistoryServiceGetSystemList;
  static readonly UpdateSystem: SapHistoryServiceUpdateSystem;
  static readonly DeleteSystem: SapHistoryServiceDeleteSystem;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class SapHistoryServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  getSapStatus(
    requestMessage: internal_saphistory_pb.SapSelection,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.TransactionData|null) => void
  ): UnaryResponse;
  getSapStatus(
    requestMessage: internal_saphistory_pb.SapSelection,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.TransactionData|null) => void
  ): UnaryResponse;
  getSystemList(
    requestMessage: internal_saphistory_pb.Nothing,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.SystemList|null) => void
  ): UnaryResponse;
  getSystemList(
    requestMessage: internal_saphistory_pb.Nothing,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.SystemList|null) => void
  ): UnaryResponse;
  updateSystem(
    requestMessage: internal_saphistory_pb.SystemInfo,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.Nothing|null) => void
  ): UnaryResponse;
  updateSystem(
    requestMessage: internal_saphistory_pb.SystemInfo,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.Nothing|null) => void
  ): UnaryResponse;
  deleteSystem(
    requestMessage: internal_saphistory_pb.SystemInfo,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.Nothing|null) => void
  ): UnaryResponse;
  deleteSystem(
    requestMessage: internal_saphistory_pb.SystemInfo,
    callback: (error: ServiceError|null, responseMessage: internal_saphistory_pb.Nothing|null) => void
  ): UnaryResponse;
}

