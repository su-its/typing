import React from "react";

interface ProgressBarProps {
  maxWidth: number;
  height: number;
  maxValue: number;
  value: number;
}

const ProgressBar: React.FC<ProgressBarProps> = ({ maxWidth, value, maxValue, height }) => {
  const rate = value / maxValue;
  const actualWidth = maxWidth * rate + "px";
  const actualHeight = height + "px";
  return <div style={{ background: "#2196f3", height: actualHeight, width: actualWidth }} />;
};

export default ProgressBar;
