/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var gogoproto_gogo_pb = require('./gogoproto/gogo_pb.js');
goog.exportSymbol('proto.swagchat.protobuf.AddUserRolesRequest', null, global);
goog.exportSymbol('proto.swagchat.protobuf.DeleteUserRolesRequest', null, global);
goog.exportSymbol('proto.swagchat.protobuf.UserRole', null, global);

/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.swagchat.protobuf.UserRole = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.swagchat.protobuf.UserRole, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.swagchat.protobuf.UserRole.displayName = 'proto.swagchat.protobuf.UserRole';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.swagchat.protobuf.UserRole.prototype.toObject = function(opt_includeInstance) {
  return proto.swagchat.protobuf.UserRole.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.swagchat.protobuf.UserRole} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.UserRole.toObject = function(includeInstance, msg) {
  var f, obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 11, ""),
    role: jspb.Message.getFieldWithDefault(msg, 12, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.swagchat.protobuf.UserRole}
 */
proto.swagchat.protobuf.UserRole.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.swagchat.protobuf.UserRole;
  return proto.swagchat.protobuf.UserRole.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.swagchat.protobuf.UserRole} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.swagchat.protobuf.UserRole}
 */
proto.swagchat.protobuf.UserRole.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 12:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRole(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.swagchat.protobuf.UserRole.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.swagchat.protobuf.UserRole.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.swagchat.protobuf.UserRole} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.UserRole.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getRole();
  if (f !== 0) {
    writer.writeInt32(
      12,
      f
    );
  }
};


/**
 * optional string user_id = 11;
 * @return {string}
 */
proto.swagchat.protobuf.UserRole.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.swagchat.protobuf.UserRole.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * optional int32 role = 12;
 * @return {number}
 */
proto.swagchat.protobuf.UserRole.prototype.getRole = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 12, 0));
};


/** @param {number} value */
proto.swagchat.protobuf.UserRole.prototype.setRole = function(value) {
  jspb.Message.setProto3IntField(this, 12, value);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.swagchat.protobuf.AddUserRolesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.swagchat.protobuf.AddUserRolesRequest.repeatedFields_, null);
};
goog.inherits(proto.swagchat.protobuf.AddUserRolesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.swagchat.protobuf.AddUserRolesRequest.displayName = 'proto.swagchat.protobuf.AddUserRolesRequest';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.swagchat.protobuf.AddUserRolesRequest.repeatedFields_ = [12];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.swagchat.protobuf.AddUserRolesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.swagchat.protobuf.AddUserRolesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.AddUserRolesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 11, ""),
    rolesList: jspb.Message.getRepeatedField(msg, 12)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.swagchat.protobuf.AddUserRolesRequest}
 */
proto.swagchat.protobuf.AddUserRolesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.swagchat.protobuf.AddUserRolesRequest;
  return proto.swagchat.protobuf.AddUserRolesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.swagchat.protobuf.AddUserRolesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.swagchat.protobuf.AddUserRolesRequest}
 */
proto.swagchat.protobuf.AddUserRolesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 12:
      var value = /** @type {!Array.<number>} */ (reader.readPackedInt32());
      msg.setRolesList(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.swagchat.protobuf.AddUserRolesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.swagchat.protobuf.AddUserRolesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.AddUserRolesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getRolesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      12,
      f
    );
  }
};


/**
 * optional string user_id = 11;
 * @return {string}
 */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * repeated int32 roles = 12;
 * @return {!Array.<number>}
 */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.getRolesList = function() {
  return /** @type {!Array.<number>} */ (jspb.Message.getRepeatedField(this, 12));
};


/** @param {!Array.<number>} value */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.setRolesList = function(value) {
  jspb.Message.setField(this, 12, value || []);
};


/**
 * @param {!number} value
 * @param {number=} opt_index
 */
proto.swagchat.protobuf.AddUserRolesRequest.prototype.addRoles = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 12, value, opt_index);
};


proto.swagchat.protobuf.AddUserRolesRequest.prototype.clearRolesList = function() {
  this.setRolesList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.swagchat.protobuf.DeleteUserRolesRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.swagchat.protobuf.DeleteUserRolesRequest.repeatedFields_, null);
};
goog.inherits(proto.swagchat.protobuf.DeleteUserRolesRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.swagchat.protobuf.DeleteUserRolesRequest.displayName = 'proto.swagchat.protobuf.DeleteUserRolesRequest';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.repeatedFields_ = [12];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.swagchat.protobuf.DeleteUserRolesRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.swagchat.protobuf.DeleteUserRolesRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    userId: jspb.Message.getFieldWithDefault(msg, 11, ""),
    rolesList: jspb.Message.getRepeatedField(msg, 12)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.swagchat.protobuf.DeleteUserRolesRequest}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.swagchat.protobuf.DeleteUserRolesRequest;
  return proto.swagchat.protobuf.DeleteUserRolesRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.swagchat.protobuf.DeleteUserRolesRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.swagchat.protobuf.DeleteUserRolesRequest}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 11:
      var value = /** @type {string} */ (reader.readString());
      msg.setUserId(value);
      break;
    case 12:
      var value = /** @type {!Array.<number>} */ (reader.readPackedInt32());
      msg.setRolesList(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.swagchat.protobuf.DeleteUserRolesRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.swagchat.protobuf.DeleteUserRolesRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getUserId();
  if (f.length > 0) {
    writer.writeString(
      11,
      f
    );
  }
  f = message.getRolesList();
  if (f.length > 0) {
    writer.writePackedInt32(
      12,
      f
    );
  }
};


/**
 * optional string user_id = 11;
 * @return {string}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.getUserId = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 11, ""));
};


/** @param {string} value */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.setUserId = function(value) {
  jspb.Message.setProto3StringField(this, 11, value);
};


/**
 * repeated int32 roles = 12;
 * @return {!Array.<number>}
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.getRolesList = function() {
  return /** @type {!Array.<number>} */ (jspb.Message.getRepeatedField(this, 12));
};


/** @param {!Array.<number>} value */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.setRolesList = function(value) {
  jspb.Message.setField(this, 12, value || []);
};


/**
 * @param {!number} value
 * @param {number=} opt_index
 */
proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.addRoles = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 12, value, opt_index);
};


proto.swagchat.protobuf.DeleteUserRolesRequest.prototype.clearRolesList = function() {
  this.setRolesList([]);
};


goog.object.extend(exports, proto.swagchat.protobuf);
