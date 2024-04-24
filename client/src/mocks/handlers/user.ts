import { delay, http, HttpResponse } from "msw";

export const MockUsers = [
  {
    id: 1,
    name: "Asa",
    image_url: "https://github.com/a01sa01to.png",
  },
  {
    id: 2,
    name: "Sor4chi",
    image_url: "https://github.com/sor4chi.png",
  },
];
export const MockMe = MockUsers[0];

export const userHandlers = [
  http.get("/api/user/me", async () => {
    await delay();
    return HttpResponse.json(MockMe);
  }),
];
