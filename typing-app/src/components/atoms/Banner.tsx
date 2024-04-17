import React from "react";
import { Box, Image, Link } from "@chakra-ui/react";
import bannerImage from "@/assets/images/banner.png";

const Banner: React.FC = () => {
  return (
    <Box>
      <Link href="/">
        <Image src={bannerImage.src} alt="Logo" maxH={54} ml={2} />
      </Link>
    </Box>
  );
};

export default Banner;
