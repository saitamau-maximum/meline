import { styles } from "./wrapper.css";

type Props = {
  children: React.ReactNode;
};

export const Wrapper = ({ children }: Props) => {
  return <div className={styles.container}>{children}</div>;
};
