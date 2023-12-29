import { http, HttpResponse } from "msw";
import { MockUsers } from "./user";

export const MockChannels = [
  {
    id: 1,
    name: "全体チャンネル",
    users: MockUsers.map((mu) => ({
      name: mu.name,
      image_url: mu.image_url,
    })),
  },
  {
    id: 2,
    name: "開発部",
    users: MockUsers.map((mu) => ({
      name: mu.name,
      image_url: mu.image_url,
    })),
  },
];

export const channelHandlers = [
  http.get("/api/channels", () =>
    HttpResponse.json({
      channels: MockChannels.map((mc) => ({
        id: mc.id,
        name: mc.name,
      })),
    })
  ),
  http.get("/api/channels/:id", ({ params }) => {
    const id = Number(params.id);
    const channel = MockChannels.find((mc) => mc.id === id);
    if (!channel) {
      return new HttpResponse(null, {
        status: 404,
      });
    }
    return HttpResponse.json({
      channel: {
        name: channel.name,
        users: channel.users,
      },
    });
  }),
];
