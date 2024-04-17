import React from "react";
import { Box, Image } from "@chakra-ui/react";
import brandImage from "@/assets/images/brand.png";

const BrandText: React.FC = () => {
  return (
    <Box>
      <Image src={brandImage.src} alt="Brand" ml={2} />
    </Box>
  );
};

export default BrandText;
