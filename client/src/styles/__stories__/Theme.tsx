import "../global.css";
import { vars } from "..";

const styles = (full: boolean) => ({
  background: vars.color.gray[2],
  color: vars.color.gray[12],
  padding: full ? "0" : "16px",
});

const LightWrapper = ({
  children,
  full = false,
}: {
  children: React.ReactNode;
  full?: boolean;
}) => (
  <div className="light" style={styles(full)}>
    {children}
  </div>
);
const DarkWrapper = ({
  children,
  full = false,
}: {
  children: React.ReactNode;
  full?: boolean;
}) => (
  <div className="dark" style={styles(full)}>
    {children}
  </div>
);

export const Theme = {
  Light: LightWrapper,
  Dark: DarkWrapper,
};
