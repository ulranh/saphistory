/**
 * @fileoverview gRPC-Web generated client stub for private
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.private = require('./saphistory_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.private.SapHistoryServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.private.SapHistoryServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.private.SapSelection,
 *   !proto.private.TransactionData>}
 */
const methodDescriptor_SapHistoryService_GetSapStatus = new grpc.web.MethodDescriptor(
  '/private.SapHistoryService/GetSapStatus',
  grpc.web.MethodType.UNARY,
  proto.private.SapSelection,
  proto.private.TransactionData,
  /**
   * @param {!proto.private.SapSelection} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.TransactionData.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.private.SapSelection,
 *   !proto.private.TransactionData>}
 */
const methodInfo_SapHistoryService_GetSapStatus = new grpc.web.AbstractClientBase.MethodInfo(
  proto.private.TransactionData,
  /**
   * @param {!proto.private.SapSelection} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.TransactionData.deserializeBinary
);


/**
 * @param {!proto.private.SapSelection} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.private.TransactionData)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.private.TransactionData>|undefined}
 *     The XHR Node Readable Stream
 */
proto.private.SapHistoryServiceClient.prototype.getSapStatus =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/private.SapHistoryService/GetSapStatus',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_GetSapStatus,
      callback);
};


/**
 * @param {!proto.private.SapSelection} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.private.TransactionData>}
 *     Promise that resolves to the response
 */
proto.private.SapHistoryServicePromiseClient.prototype.getSapStatus =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/private.SapHistoryService/GetSapStatus',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_GetSapStatus);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.private.Nothing,
 *   !proto.private.SystemList>}
 */
const methodDescriptor_SapHistoryService_GetSystemList = new grpc.web.MethodDescriptor(
  '/private.SapHistoryService/GetSystemList',
  grpc.web.MethodType.UNARY,
  proto.private.Nothing,
  proto.private.SystemList,
  /**
   * @param {!proto.private.Nothing} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.SystemList.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.private.Nothing,
 *   !proto.private.SystemList>}
 */
const methodInfo_SapHistoryService_GetSystemList = new grpc.web.AbstractClientBase.MethodInfo(
  proto.private.SystemList,
  /**
   * @param {!proto.private.Nothing} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.SystemList.deserializeBinary
);


/**
 * @param {!proto.private.Nothing} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.private.SystemList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.private.SystemList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.private.SapHistoryServiceClient.prototype.getSystemList =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/private.SapHistoryService/GetSystemList',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_GetSystemList,
      callback);
};


/**
 * @param {!proto.private.Nothing} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.private.SystemList>}
 *     Promise that resolves to the response
 */
proto.private.SapHistoryServicePromiseClient.prototype.getSystemList =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/private.SapHistoryService/GetSystemList',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_GetSystemList);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.private.SystemInfo,
 *   !proto.private.Nothing>}
 */
const methodDescriptor_SapHistoryService_UpdateSystem = new grpc.web.MethodDescriptor(
  '/private.SapHistoryService/UpdateSystem',
  grpc.web.MethodType.UNARY,
  proto.private.SystemInfo,
  proto.private.Nothing,
  /**
   * @param {!proto.private.SystemInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.Nothing.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.private.SystemInfo,
 *   !proto.private.Nothing>}
 */
const methodInfo_SapHistoryService_UpdateSystem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.private.Nothing,
  /**
   * @param {!proto.private.SystemInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.Nothing.deserializeBinary
);


/**
 * @param {!proto.private.SystemInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.private.Nothing)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.private.Nothing>|undefined}
 *     The XHR Node Readable Stream
 */
proto.private.SapHistoryServiceClient.prototype.updateSystem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/private.SapHistoryService/UpdateSystem',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_UpdateSystem,
      callback);
};


/**
 * @param {!proto.private.SystemInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.private.Nothing>}
 *     Promise that resolves to the response
 */
proto.private.SapHistoryServicePromiseClient.prototype.updateSystem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/private.SapHistoryService/UpdateSystem',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_UpdateSystem);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.private.SystemInfo,
 *   !proto.private.Nothing>}
 */
const methodDescriptor_SapHistoryService_DeleteSystem = new grpc.web.MethodDescriptor(
  '/private.SapHistoryService/DeleteSystem',
  grpc.web.MethodType.UNARY,
  proto.private.SystemInfo,
  proto.private.Nothing,
  /**
   * @param {!proto.private.SystemInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.Nothing.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.private.SystemInfo,
 *   !proto.private.Nothing>}
 */
const methodInfo_SapHistoryService_DeleteSystem = new grpc.web.AbstractClientBase.MethodInfo(
  proto.private.Nothing,
  /**
   * @param {!proto.private.SystemInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.private.Nothing.deserializeBinary
);


/**
 * @param {!proto.private.SystemInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.private.Nothing)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.private.Nothing>|undefined}
 *     The XHR Node Readable Stream
 */
proto.private.SapHistoryServiceClient.prototype.deleteSystem =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/private.SapHistoryService/DeleteSystem',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_DeleteSystem,
      callback);
};


/**
 * @param {!proto.private.SystemInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.private.Nothing>}
 *     Promise that resolves to the response
 */
proto.private.SapHistoryServicePromiseClient.prototype.deleteSystem =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/private.SapHistoryService/DeleteSystem',
      request,
      metadata || {},
      methodDescriptor_SapHistoryService_DeleteSystem);
};


module.exports = proto.private;

