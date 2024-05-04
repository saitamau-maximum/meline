import { useEffect, useRef } from "react";
import { styles } from "./chat-display.css";
import { Message } from "@/repositories/message";

interface ChatDisplayProps {
  messages: Message[];
}

export const ChatDisplay = ({ messages }: ChatDisplayProps) => {
  const chatDisplayRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const chatDisplay = chatDisplayRef.current;
    chatDisplay?.scrollTo({ top: chatDisplay.scrollHeight });
  }, [messages]);

  return (
    <div ref={chatDisplayRef} className={styles.container}>
      {messages.map((message) => (
        <div key={message.id} className={styles.messageContainer}>
          {message.reply_to_message && (
            <div className={styles.replyMessage}>
              <img
                className={styles.replyAvatar}
                src={message.reply_to_message.user.image_url}
                alt={message.reply_to_message.user.name}
              />
              <div className={styles.replyAuthor}>
                {message.reply_to_message.user.name}
              </div>
              <div className={styles.replyText}>
                {message.reply_to_message.content}
              </div>
            </div>
          )}
          <div className={styles.message}>
            <img
              className={styles.avatar}
              src={message.user.image_url}
              alt={message.user.name}
            />
            <div className={styles.messageContent}>
              <div className={styles.messageAuthor}>{message.user.name}</div>
              <div className={styles.messageText}>{message.content}</div>
            </div>
          </div>
        </div>
      ))}
    </div>
  );
};
