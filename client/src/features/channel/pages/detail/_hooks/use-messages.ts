import { useQuery } from "@tanstack/react-query";
import { Message, MessageRepositoryImpl } from "@/repositories/message";
import { useMemo, useState } from "react";

interface UseMessagesOptions {
  channelId: number;
}

export const useMessages = ({ channelId }: UseMessagesOptions) => {
  const query = useQuery({
    queryKey: MessageRepositoryImpl.getMessages$$key(channelId),
    queryFn: () => MessageRepositoryImpl.getMessages(channelId),
  });

  const [postedMessages, setPostedMessages] = useState<Message[]>([]);

  const appendMessage = (message: Message) => {
    setPostedMessages((prev) => [...prev, message]);
  };

  const messages = useMemo(
    () => [...(query.data?.messages ?? []), ...postedMessages],
    [query.data?.messages, postedMessages]
  );

  return {
    messages,
    appendMessage,
  };
};
