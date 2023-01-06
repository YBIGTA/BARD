import { CloseIcon } from '@chakra-ui/icons';
import { Box, Flex, Image, Spinner, Text } from '@chakra-ui/react';
import { forwardRef } from 'react';

const PhotoInputItem = forwardRef(
  ({ src, isUploading, onDelete, index }, ref) => {
    return (
      <>
        <Box
          h="80px"
          w="80px"
          position="relative"
          overflow="visible"
          boxShadow={'md'}
          borderBottomLeftRadius="5px"
          flexShrink={0}
          ref={ref}
        >
          {isUploading && (
            <>
              <Box
                h="80px"
                w="80px"
                background="black"
                opacity="0.5"
                position="absolute"
                top="0"
                left="0"
                borderRadius="5px"
              />
              <Spinner
                position="absolute"
                left="28px"
                top="28px"
                color="yellow.500"
              />
            </>
          )}
          {!isUploading && (
            <Flex
              w="12px"
              h="12px"
              rounded="full"
              position="absolute"
              right="-4px"
              top="-4px"
              zIndex="5"
              background="white"
              justify="center"
              align="center"
            >
              <CloseIcon fontSize="6px" strokeWidth="1.5" onClick={onDelete} />
            </Flex>
          )}
          <Box
            w="0"
            h="0"
            position="absolute"
            borderLeft="20px solid"
            borderLeftColor="yellow.200"
            borderBottom="20px solid"
            borderBottomColor="yellow.200"
            borderRight="20px solid"
            borderRightColor="transparent"
            borderTop="20px solid"
            borderTopColor="transparent"
            bottom="0px"
            left="0px"
            isolation="isolate"
            borderBottomLeftRadius="5px"
            pointerEvents="none"
          >
            <Text
              size="3xs"
              position="absolute"
              left="-16px"
              bottom="-22px"
            >{`#${index + 1}`}</Text>
          </Box>
          <Image
            h="100%"
            w="100%"
            objectFit="cover"
            src={src}
            borderRadius="5px"
          />
        </Box>
      </>
    );
  }
);

export default PhotoInputItem;
