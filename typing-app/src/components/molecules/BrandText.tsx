import React from "react";
import { Box, HStack, Text } from "@chakra-ui/react";

const BrandText: React.FC = () => {
  return (
    <Box>
      <HStack>
        <Text fontSize="50px" color={"white"}>
          ITS
        </Text>
        <Text fontSize="20px" color={"white"}>
          静岡大学ITソルーション室
        </Text>
      </HStack>
    </Box>
  );
};

export default BrandText;
