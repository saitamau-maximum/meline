import { delay, http, HttpResponse } from "msw";
import { MockMe, MockUsers } from "./user";

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
  http.get("/api/channel", async () => {
    await delay();
    return HttpResponse.json({
      channels: MockChannels.map((mc) => ({
        id: mc.id,
        name: mc.name,
      })),
    });
  }),
  http.post("/api/channel", async ({ request }) => {
    const body = (await request.json()) as { name: string };
    MockChannels.push({
      id: MockChannels.length + 1,
      name: body.name,
      users: [MockMe],
    });
    await delay();
    return new HttpResponse(null);
  }),
  http.get(
    "/api/channel/:id",
    async ({ params }) => {
      await delay();
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
    },
    {}
  ),
];
