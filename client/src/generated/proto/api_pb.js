// @generated by protoc-gen-es v2.2.2
// @generated from file api.proto (package meline, syntax proto3)
/* eslint-disable */

import { fileDesc, serviceDesc } from "@bufbuild/protobuf/codegenv1";
import { file_google_api_annotations } from "./google/api/annotations_pb";
import { file_channel_request } from "./channel_request_pb";
import { file_channel_response } from "./channel_response_pb";
import { file_message_request } from "./message_request_pb";
import { file_message_response } from "./message_response_pb";
import { file_user_request } from "./user_request_pb";
import { file_user_response } from "./user_response_pb";
import { file_notify_request } from "./notify_request_pb";
import { file_notify_response } from "./notify_response_pb";

/**
 * Describes the file api.proto.
 */
export const file_api = /*@__PURE__*/
  fileDesc("CglhcGkucHJvdG8SBm1lbGluZTLsBwoOQ2hhbm5lbFNlcnZpY2UScwoOR2V0QWxsQ2hhbm5lbHMSJS5zY2hlbWEucmVxdWVzdC5HZXRBbGxDaGFubmVsc1JlcXVlc3QaJy5zY2hlbWEucmVzcG9uc2UuR2V0QWxsQ2hhbm5lbHNSZXNwb25zZSIRgtPkkwILEgkvY2hhbm5lbHMSeAoOR2V0Q2hhbm5lbEJ5SUQSJS5zY2hlbWEucmVxdWVzdC5HZXRDaGFubmVsQnlJRFJlcXVlc3QaJy5zY2hlbWEucmVzcG9uc2UuR2V0Q2hhbm5lbEJ5SURSZXNwb25zZSIWgtPkkwIQEg4vY2hhbm5lbHMve2lkfRJ3CgtKb2luQ2hhbm5lbBIiLnNjaGVtYS5yZXF1ZXN0LkpvaW5DaGFubmVsUmVxdWVzdBokLnNjaGVtYS5yZXNwb25zZS5Kb2luQ2hhbm5lbFJlc3BvbnNlIh6C0+STAhg6ASoiEy9jaGFubmVscy97aWR9L2pvaW4ScwoNQ3JlYXRlQ2hhbm5lbBIkLnNjaGVtYS5yZXF1ZXN0LkNyZWF0ZUNoYW5uZWxSZXF1ZXN0GiYuc2NoZW1hLnJlc3BvbnNlLkNyZWF0ZUNoYW5uZWxSZXNwb25zZSIUgtPkkwIOOgEqIgkvY2hhbm5lbHMSjgEKEkNyZWF0ZUNoaWxkQ2hhbm5lbBIpLnNjaGVtYS5yZXF1ZXN0LkNyZWF0ZUNoaWxkQ2hhbm5lbFJlcXVlc3QaKy5zY2hlbWEucmVzcG9uc2UuQ3JlYXRlQ2hpbGRDaGFubmVsUmVzcG9uc2UiIILT5JMCGjoBKiIVL2NoYW5uZWxzL3tpZH0vY3JlYXRlEngKDVVwZGF0ZUNoYW5uZWwSJC5zY2hlbWEucmVxdWVzdC5VcGRhdGVDaGFubmVsUmVxdWVzdBomLnNjaGVtYS5yZXNwb25zZS5VcGRhdGVDaGFubmVsUmVzcG9uc2UiGYLT5JMCEzoBKhoOL2NoYW5uZWxzL3tpZH0SdQoNRGVsZXRlQ2hhbm5lbBIkLnNjaGVtYS5yZXF1ZXN0LkRlbGV0ZUNoYW5uZWxSZXF1ZXN0GiYuc2NoZW1hLnJlc3BvbnNlLkRlbGV0ZUNoYW5uZWxSZXNwb25zZSIWgtPkkwIQKg4vY2hhbm5lbHMve2lkfRJ7CgxMZWF2ZUNoYW5uZWwSIy5zY2hlbWEucmVxdWVzdC5MZWF2ZUNoYW5uZWxSZXF1ZXN0GiUuc2NoZW1hLnJlc3BvbnNlLkxlYXZlQ2hhbm5lbFJlc3BvbnNlIh+C0+STAhk6ASoiFC9jaGFubmVscy97aWR9L2xlYXZlMqEFCg5NZXNzYWdlU2VydmljZRKAAQoOR2V0QnlDaGFubmVsSUQSJS5zY2hlbWEucmVxdWVzdC5HZXRCeUNoYW5uZWxJRFJlcXVlc3QaJy5zY2hlbWEucmVzcG9uc2UuR2V0QnlDaGFubmVsSURSZXNwb25zZSIegtPkkwIYEhYve2NoYW5uZWxfaWR9L21lc3NhZ2VzEnkKBkNyZWF0ZRIkLnNjaGVtYS5yZXF1ZXN0LkNyZWF0ZU1lc3NhZ2VSZXF1ZXN0GiYuc2NoZW1hLnJlc3BvbnNlLkNyZWF0ZU1lc3NhZ2VSZXNwb25zZSIhgtPkkwIbOgEqIhYve2NoYW5uZWxfaWR9L21lc3NhZ2VzEpMBCgtDcmVhdGVSZXBseRIpLnNjaGVtYS5yZXF1ZXN0LkNyZWF0ZVJlcGx5TWVzc2FnZVJlcXVlc3QaKy5zY2hlbWEucmVzcG9uc2UuQ3JlYXRlUmVwbHlNZXNzYWdlUmVzcG9uc2UiLILT5JMCJjoBKiIhL3tjaGFubmVsX2lkfS9tZXNzYWdlcy97aWR9L3JlcGx5En4KBlVwZGF0ZRIkLnNjaGVtYS5yZXF1ZXN0LlVwZGF0ZU1lc3NhZ2VSZXF1ZXN0GiYuc2NoZW1hLnJlc3BvbnNlLlVwZGF0ZU1lc3NhZ2VSZXNwb25zZSImgtPkkwIgOgEqGhsve2NoYW5uZWxfaWR9L21lc3NhZ2VzL3tpZH0SewoGRGVsZXRlEiQuc2NoZW1hLnJlcXVlc3QuRGVsZXRlTWVzc2FnZVJlcXVlc3QaJi5zY2hlbWEucmVzcG9uc2UuRGVsZXRlTWVzc2FnZVJlc3BvbnNlIiOC0+STAh0qGy97Y2hhbm5lbF9pZH0vbWVzc2FnZXMve2lkfTJmCgtVc2VyU2VydmljZRJXCgJNZRIdLnNjaGVtYS5yZXF1ZXN0LlVzZXJNZVJlcXVlc3QaHy5zY2hlbWEucmVzcG9uc2UuVXNlck1lUmVzcG9uc2UiEYLT5JMCCxIJL3VzZXJzL21lMnAKDU5vdGlmeVNlcnZpY2USXwoGTm90aWZ5Eh0uc2NoZW1hLnJlcXVlc3QuTm90aWZ5UmVxdWVzdBofLnNjaGVtYS5yZXNwb25zZS5Ob3RpZnlSZXNwb25zZSIVgtPkkwIPOgEqIgovd3Mvbm90aWZ5QngKCmNvbS5tZWxpbmVCCEFwaVByb3RvUAFaKGdpdGh1Yi5jb20vc2FpdGFtYXUtbWF4aW11bS9tZWxpbmUvcHJvdG+iAgNNWFiqAgZNZWxpbmXKAgZNZWxpbmXiAhJNZWxpbmVcR1BCTWV0YWRhdGHqAgZNZWxpbmViBnByb3RvMw", [file_google_api_annotations, file_channel_request, file_channel_response, file_message_request, file_message_response, file_user_request, file_user_response, file_notify_request, file_notify_response]);

/**
 * @generated from service meline.ChannelService
 */
export const ChannelService = /*@__PURE__*/
  serviceDesc(file_api, 0);

/**
 * @generated from service meline.MessageService
 */
export const MessageService = /*@__PURE__*/
  serviceDesc(file_api, 1);

/**
 * @generated from service meline.UserService
 */
export const UserService = /*@__PURE__*/
  serviceDesc(file_api, 2);

/**
 * @generated from service meline.NotifyService
 */
export const NotifyService = /*@__PURE__*/
  serviceDesc(file_api, 3);

