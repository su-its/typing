import React from "react";

interface ProgressBarProps {
  maxWidth: number;
  value: number; // 0-100 [%]
}

const ProgressBar: React.FC<ProgressBarProps> = ({ maxWidth, value }) => {
  const width = (maxWidth * value) / 100 + "px";
  return <div style={{ background: "#2196f3", height: "20px", width: width }} />;
};

export default ProgressBar;
