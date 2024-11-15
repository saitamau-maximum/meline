import { useParams } from "react-router-dom";
import { ChannelDetailPageTemplate } from "./template";
import { useMessages } from "./_hooks/use-messages";

export const ChannelDetailPage = () => {
  const { channelId } = useParams<{ channelId: string }>();
  const { messages } = useMessages({
    channelId: Number(channelId),
  });

  return (
    <>
      <ChannelDetailPageTemplate
        channelId={Number(channelId)}
        messages={messages ?? []}
      />
    </>
  );
};
