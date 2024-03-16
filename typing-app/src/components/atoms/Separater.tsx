import React from "react";
import { Box } from "@chakra-ui/react";

interface SeparatorProps {
  height?: string;
  backgroundColor?: string;
}

const Separator: React.FC<SeparatorProps> = ({ height = "5px", backgroundColor = "white" }) => {
  return <Box height={height} bg={backgroundColor} width="100%" />;
};

export default Separator;
