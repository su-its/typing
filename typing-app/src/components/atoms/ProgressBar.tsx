import React from "react";

interface ProgressBarProps {
  maxWidth: number;
  height: number;
  value: number; // 0-100 [%]
}

const ProgressBar: React.FC<ProgressBarProps> = ({ maxWidth, value, height }) => {
  const actualWidth = (maxWidth * value) / 100 + "px";
  const actualHeight = height + "px";
  return <div style={{ background: "#2196f3", height: actualHeight, width: actualWidth }} />;
};

export default ProgressBar;
