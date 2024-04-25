import { styles } from "./chat-layout.css";

interface ChatLayoutProps {
  footer: React.ReactNode;
  children: React.ReactNode;
}

export const ChatLayout = ({ footer, children }: ChatLayoutProps) => {
  return (
    <div className={styles.chatLayout}>
      <div className={styles.chatLayoutContent}>{children}</div>
      <div className={styles.chatLayoutFooter}>{footer}</div>
    </div>
  );
};
