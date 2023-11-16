import "../global.css";
import { vars } from "../theme.css";

const styles = {
  background: vars.color.gray[1],
  color: vars.color.gray[12],
  padding: "16px",
  border: `1px solid ${vars.color.gray[6]}`,
  borderRadius: "8px",
  width: "fit-content",
};

const LightWrapper = ({ children }: { children: React.ReactNode }) => (
  <div className="light" style={styles}>
    {children}
  </div>
);

const DarkWrapper = ({ children }: { children: React.ReactNode }) => (
  <div className="dark" style={styles}>
    {children}
  </div>
);

export const Theme = {
  Light: LightWrapper,
  Dark: DarkWrapper,
};
