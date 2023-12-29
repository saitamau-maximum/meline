import { MaximumIcon } from "@/components/icons/maximum";
import { styles } from "./layout.css";

interface ChannelLayoutProps {
  sidePanel: React.ReactNode;
  main: React.ReactNode;
}

export const ChannelLayout = ({ sidePanel, main }: ChannelLayoutProps) => {
  return (
    <div className={styles.channelLayoutWrapper}>
      <div className={styles.channelLayoutSidePanel}>
        <div className={styles.channelLayoutSidePanelLogo}>
          <MaximumIcon />
        </div>
        {sidePanel}
      </div>
      <div className={styles.channelLayoutMain}>{main}</div>
    </div>
  );
};
