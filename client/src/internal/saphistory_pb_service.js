// package: private
// file: internal/saphistory.proto

var internal_saphistory_pb = require("../internal/saphistory_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var SapHistoryService = (function () {
  function SapHistoryService() {}
  SapHistoryService.serviceName = "private.SapHistoryService";
  return SapHistoryService;
}());

SapHistoryService.GetSapStatus = {
  methodName: "GetSapStatus",
  service: SapHistoryService,
  requestStream: false,
  responseStream: false,
  requestType: internal_saphistory_pb.SapSelection,
  responseType: internal_saphistory_pb.TransactionData
};

SapHistoryService.GetSystemList = {
  methodName: "GetSystemList",
  service: SapHistoryService,
  requestStream: false,
  responseStream: false,
  requestType: internal_saphistory_pb.Nothing,
  responseType: internal_saphistory_pb.SystemList
};

SapHistoryService.UpdateSystem = {
  methodName: "UpdateSystem",
  service: SapHistoryService,
  requestStream: false,
  responseStream: false,
  requestType: internal_saphistory_pb.SystemInfo,
  responseType: internal_saphistory_pb.Nothing
};

SapHistoryService.DeleteSystem = {
  methodName: "DeleteSystem",
  service: SapHistoryService,
  requestStream: false,
  responseStream: false,
  requestType: internal_saphistory_pb.SystemInfo,
  responseType: internal_saphistory_pb.Nothing
};

exports.SapHistoryService = SapHistoryService;

function SapHistoryServiceClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

SapHistoryServiceClient.prototype.getSapStatus = function getSapStatus(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SapHistoryService.GetSapStatus, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SapHistoryServiceClient.prototype.getSystemList = function getSystemList(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SapHistoryService.GetSystemList, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SapHistoryServiceClient.prototype.updateSystem = function updateSystem(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SapHistoryService.UpdateSystem, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

SapHistoryServiceClient.prototype.deleteSystem = function deleteSystem(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(SapHistoryService.DeleteSystem, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.SapHistoryServiceClient = SapHistoryServiceClient;

