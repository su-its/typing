// src/components/atoms/FooterBrand.tsx
import React from "react";
import { Box, HStack, Text } from "@chakra-ui/react";

const BrandText: React.FC = () => {
  return (
    <Box>
      <HStack>
        <Text fontSize="50px">ITS</Text>
        <Text fontSize="sm">静岡大学ITソルーション室</Text>
      </HStack>
    </Box>
  );
};

export default BrandText;
