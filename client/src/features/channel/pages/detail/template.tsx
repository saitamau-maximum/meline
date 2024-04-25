import { Message } from "@/repositories/message";
import { ChatDisplay } from "./_components/chat-display";
import { ChatInput } from "./_components/chat-input/chat-input";
import { ChatLayout } from "./_components/chat-layout";

interface ChannelDetailPageProps {
  channelId: number;
  messages: Message[];
}

export const ChannelDetailPageTemplate = ({
  channelId,
  messages,
}: ChannelDetailPageProps) => {
  return (
    <ChatLayout footer={<ChatInput channelId={channelId} />}>
      <ChatDisplay messages={messages} />
    </ChatLayout>
  );
};
