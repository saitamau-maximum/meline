import { useParams } from "react-router-dom";
import { ChannelDetailPageTemplate } from "./template";
import { useMessages } from "./_hooks/use-messages";
import { useChat } from "./_hooks/use-chat";

export const ChannelDetailPage = () => {
  const { channelId } = useParams<{ channelId: string }>();
  const { messages, appendMessage } = useMessages({
    channelId: Number(channelId),
  });
  useChat({
    channelId: Number(channelId),
    onMessageReceived: (res) => {
      appendMessage(res.message);
    },
  });

  return (
    <ChannelDetailPageTemplate
      channelId={Number(channelId)}
      messages={messages ?? []}
    />
  );
};
