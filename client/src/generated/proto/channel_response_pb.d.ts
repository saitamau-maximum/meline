// @generated by protoc-gen-es v2.2.2
// @generated from file channel_response.proto (package schema.response, syntax proto3)
/* eslint-disable */

import type { GenFile, GenMessage } from "@bufbuild/protobuf/codegenv1";
import type { Message } from "@bufbuild/protobuf";
import type { Channel, ChannelDetail } from "./channel_base_pb";

/**
 * Describes the file channel_response.proto.
 */
export declare const file_channel_response: GenFile;

/**
 * @generated from message schema.response.GetAllChannelsResponse
 */
export declare type GetAllChannelsResponse = Message<"schema.response.GetAllChannelsResponse"> & {
  /**
   * @generated from field: repeated base.Channel channels = 1;
   */
  channels: Channel[];
};

/**
 * Describes the message schema.response.GetAllChannelsResponse.
 * Use `create(GetAllChannelsResponseSchema)` to create a new message.
 */
export declare const GetAllChannelsResponseSchema: GenMessage<GetAllChannelsResponse>;

/**
 * @generated from message schema.response.GetChannelByIDResponse
 */
export declare type GetChannelByIDResponse = Message<"schema.response.GetChannelByIDResponse"> & {
  /**
   * @generated from field: base.ChannelDetail channel = 1;
   */
  channel?: ChannelDetail;
};

/**
 * Describes the message schema.response.GetChannelByIDResponse.
 * Use `create(GetChannelByIDResponseSchema)` to create a new message.
 */
export declare const GetChannelByIDResponseSchema: GenMessage<GetChannelByIDResponse>;

/**
 * @generated from message schema.response.JoinChannelResponse
 */
export declare type JoinChannelResponse = Message<"schema.response.JoinChannelResponse"> & {
};

/**
 * Describes the message schema.response.JoinChannelResponse.
 * Use `create(JoinChannelResponseSchema)` to create a new message.
 */
export declare const JoinChannelResponseSchema: GenMessage<JoinChannelResponse>;

/**
 * @generated from message schema.response.CreateChannelResponse
 */
export declare type CreateChannelResponse = Message<"schema.response.CreateChannelResponse"> & {
};

/**
 * Describes the message schema.response.CreateChannelResponse.
 * Use `create(CreateChannelResponseSchema)` to create a new message.
 */
export declare const CreateChannelResponseSchema: GenMessage<CreateChannelResponse>;

/**
 * @generated from message schema.response.CreateChildChannelResponse
 */
export declare type CreateChildChannelResponse = Message<"schema.response.CreateChildChannelResponse"> & {
};

/**
 * Describes the message schema.response.CreateChildChannelResponse.
 * Use `create(CreateChildChannelResponseSchema)` to create a new message.
 */
export declare const CreateChildChannelResponseSchema: GenMessage<CreateChildChannelResponse>;

/**
 * @generated from message schema.response.UpdateChannelResponse
 */
export declare type UpdateChannelResponse = Message<"schema.response.UpdateChannelResponse"> & {
};

/**
 * Describes the message schema.response.UpdateChannelResponse.
 * Use `create(UpdateChannelResponseSchema)` to create a new message.
 */
export declare const UpdateChannelResponseSchema: GenMessage<UpdateChannelResponse>;

/**
 * @generated from message schema.response.DeleteChannelResponse
 */
export declare type DeleteChannelResponse = Message<"schema.response.DeleteChannelResponse"> & {
};

/**
 * Describes the message schema.response.DeleteChannelResponse.
 * Use `create(DeleteChannelResponseSchema)` to create a new message.
 */
export declare const DeleteChannelResponseSchema: GenMessage<DeleteChannelResponse>;

/**
 * @generated from message schema.response.LeaveChannelResponse
 */
export declare type LeaveChannelResponse = Message<"schema.response.LeaveChannelResponse"> & {
};

/**
 * Describes the message schema.response.LeaveChannelResponse.
 * Use `create(LeaveChannelResponseSchema)` to create a new message.
 */
export declare const LeaveChannelResponseSchema: GenMessage<LeaveChannelResponse>;

