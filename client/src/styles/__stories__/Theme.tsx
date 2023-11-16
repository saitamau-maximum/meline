import "../global.css";

const LightWrapper = ({ children }: { children: React.ReactNode }) => (
    <div className="light">{children}</div>
);

const DarkWrapper = ({ children }: { children: React.ReactNode }) => (
    <div className="dark">{children}</div>
);

export const Theme = {
    Light: LightWrapper,
    Dark: DarkWrapper,
};
