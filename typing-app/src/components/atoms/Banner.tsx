import React from "react";
import { Box, Image } from "@chakra-ui/react";
import bannerImage from "@/assets/images/banner.png";

const Banner: React.FC = () => {
  return (
    <Box>
      <Image src={bannerImage.src} alt="Logo" maxH={100} />
    </Box>
  );
};

export default Banner;
