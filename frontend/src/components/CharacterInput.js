import { CheckIcon, RepeatIcon } from '@chakra-ui/icons';
import {
  Box,
  Text,
  Flex,
  Button,
  Input,
  TagLabel,
  Tag,
  TagCloseButton,
} from '@chakra-ui/react';
import { useState } from 'react';

// 🥷 닌자, 🧞 지니, 🧟 좀비, 🧑‍🚀 우주비행사, 🐇 달토끼, 🧙‍♂️ 간달프, 🧛 드라큘라, 🦈 아기상어, ⛄ 눈사람, 🦅 아기독수리, 🦔 소닉, 🦑 오징어, 🧜‍♀️ 에리얼, 🧚‍♀️ 팅커벨, 🦸‍♂️ 슈퍼맨, 🦹‍♂️ 배트맨, 🗡️ 전사, 🌝 토마스, 🚖 범블비, 🌞 햇님, 🌛 달님, 🍄 마리오, 🌳 그루트, 🕷️ 스파이더맨, 🐜 앤트맨, ☃️ 올라프, 🧝‍♂️ 레골라스, 🧪 매드 사이언티스트

const CharacterInput = ({ characters, setCharacters }) => {
  const [suggestions, setSuggestions] = useState([
    {
      emoji: '🥷',
      name: '닌자',
    },
    {
      emoji: '🧞',
      name: '지니',
    },
    {
      emoji: '🧟',
      name: '좀비',
    },
    {
      emoji: '🧑‍🚀',
      name: '우주비행사',
    },
    {
      emoji: '🐇',
      name: '달토끼',
    },
    {
      emoji: '🧙‍♂️',
      name: '간달프',
    },
    {
      emoji: '🧛',
      name: '드라큘라',
    },
    {
      emoji: '🦈',
      name: '아기상어',
    },
    {
      emoji: '⛄',
      name: '눈사람',
    },
    {
      emoji: '🦅',
      name: '아기독수리',
    },
    {
      emoji: '🦔',
      name: '소닉',
    },
    {
      emoji: '🦑',
      name: '오징어',
    },
    {
      emoji: '🧜‍♀️',
      name: '에리얼',
    },
    {
      emoji: '🧚‍♀️',
      name: '팅커벨',
    },
    {
      emoji: '🦸‍♂️',
      name: '슈퍼맨',
    },
    {
      emoji: '🦹‍♂️',
      name: '배트맨',
    },
    {
      emoji: '🗡️',
      name: '전사',
    },
    {
      emoji: '🌝',
      name: '토마스',
    },
    {
      emoji: '🚖',
      name: '범블비',
    },
    {
      emoji: '🌞',
      name: '햇님',
    },
    {
      emoji: '🌛',
      name: '달님',
    },
    {
      emoji: '🍄',
      name: '마리오',
    },
    {
      emoji: '🌳',
      name: '그루트',
    },
    {
      emoji: '🕷️',
      name: '스파이더맨',
    },
    {
      emoji: '🐜',
      name: '앤트맨',
    },
    {
      emoji: '☃️',
      name: '올라프',
    },
    {
      emoji: '🧝‍♂️',
      name: '레골라스',
    },
    {
      emoji: '🧪',
      name: '매드 사이언티스트',
    },
    {
      emoji: '⚽️',
      name: '네이마르',
    },
    {
      emoji: '🎄',
      name: '산타',
    },
    {
      emoji: '🦌',
      name: '루돌프',
    },
    {
      emoji: '🐶',
      name: '강아지',
    },
    {
      emoji: '😼',
      name: '고양이',
    },
  ]);

  const [name, setName] = useState('');

  const onNameSubmit = () => {
    const targetName = name.trim();
    if (name.trim() !== '') {
      if (characters.includes(targetName)) {
        alert('이미 추가된 캐릭터입니다.');
        return;
      }
      setCharacters(prev => [...prev, targetName]);
    }
    setName('');
  };

  const onSelectSuggestion = name => {
    if (characters.includes(name)) {
      alert('이미 추가된 캐릭터입니다.');
      return;
    }
    setCharacters(prev => [...prev, name]);
  };

  const removeCharacter = name => {
    setCharacters(prev => prev.filter(item => item !== name));
  };

  const shuffleSuggestions = () => {
    console.log(suggestions);
    const arr = Array.from(suggestions);
    for (let idx = arr.length - 1; idx > 0; idx--) {
      const randInt = Math.floor(Math.random() * (idx + 1));
      const temporary = arr[idx];
      arr[idx] = arr[randInt];
      arr[randInt] = temporary;
    }
    console.log(arr);
    setSuggestions(arr);
  };

  return (
    <Flex flexDirection="column" gap="20px">
      <Flex flexDirection="column" align="flex-start" gap="5px">
        <Flex gap="5px">
          <Text fontSize="sm">* BARD의 캐릭터 제안</Text>
          <Button colorScheme="yellow" size="xs" onClick={shuffleSuggestions}>
            <RepeatIcon />
          </Button>
        </Flex>
        <Flex
          minH="48px"
          w="calc(100vw - 48px)"
          maxW="calc(768px - 48px)"
          border="1px solid"
          borderRadius="5px"
          padding="8px"
          gap="4px"
          overflow="auto"
        >
          {suggestions.slice(0, 3).map((item, index) => (
            <Button
              key={index}
              bgColor="#FDFAEF"
              flexShrink={0}
              onClick={() => onSelectSuggestion(item.name)}
            >
              <Box>{item.emoji}</Box>
              <Box>{item.name}</Box>
            </Button>
          ))}
        </Flex>
      </Flex>
      <Text fontWeight="bold">캐릭터 입력</Text>
      <Flex flexDirection="row" gap="8px">
        <Input
          value={name}
          onChange={e => setName(e.target.value)}
          placeholder="캐릭터를 입력하세요. ex: 서혁준"
        />
        <Button colorScheme="yellow" onClick={onNameSubmit}>
          <CheckIcon />
        </Button>
      </Flex>
      <Box
        w="calc(100vw - 48px)"
        maxW="calc(768px - 48px)"
        minH="100px"
        background="gray.100"
        p="10px"
      >
        {characters.map((item, index) => (
          <Tag key={index} colorScheme="blackAlpha" variant="subtle" m="0.2rem">
            <TagLabel>{item}</TagLabel>
            <TagCloseButton onClick={() => removeCharacter(item)} />
          </Tag>
        ))}
      </Box>
    </Flex>
  );
};

export default CharacterInput;
