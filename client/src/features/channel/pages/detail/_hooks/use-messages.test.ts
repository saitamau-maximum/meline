import { renderHook, waitFor } from "@/libs/test-utils";
import { IMessageRepository } from "@/repositories/message";
import { useMessages } from "./use-messages";

const mockMessageRepository = {
  createMessage: vi.fn(),
  getMessages: vi.fn(),
  getMessages$$key: vi.fn(),
} satisfies IMessageRepository;

describe("useMessages", () => {
  afterEach(() => {
    vi.clearAllMocks();
  });

  it("should return messages", async () => {
    const channelId = 1;
    const messages = [{ id: 1, content: "Hello" }];
    mockMessageRepository.getMessages.mockResolvedValue({ messages });
    mockMessageRepository.getMessages$$key.mockReturnValue([
      "messages",
      channelId,
    ]);

    const { result, rerender } = renderHook(() => useMessages({ channelId }), {
      repositories: {
        messageRepository: mockMessageRepository,
      },
    });

    await waitFor(() => expect(result.current.messages).toEqual(messages));

    rerender();
  });
});
