"use client";

import styled, { css } from "styled-components";

export const Title = styled.h2`
  margin: 20px !important;
  text-align: center;
  color: var(--blue);
`;

export const Main = styled.main`
  height: 100vh;
  overflow-y: auto;
  background-color: var(--gray-soft);
`;

export const ImageWrapper = styled.div`
  border-radius: 8px;
  position: relative;

  & img {
    border-radius: 8px 8px 2px 2px;
    position: absolute;
    width: 100%;
    height: 100%;
  }
`;

export const Divider = styled.hr`
  margin: 12px;
  background-color: var(--purple);
  height: 0;
`;

export const HR = styled.hr`
  border: 0;
  height: 2px;
  width: 40%;
`;

export const RGradientHr = styled(HR)<{ color?: string; width?: string }>`
  ${({ width }) =>
    width &&
    css`
      width: ${width};
    `};

  ${({ color }) =>
    css`
      background-image: linear-gradient(
        90deg,
        ${color ? `var(--${color})` : "var(--violete)"},
        transparent
      );
    `};
`;

export const LGradientHr = styled(HR)<{ color?: string; width?: string }>`
  ${({ width }) =>
    width &&
    css`
      width: ${width};
    `};

  ${({ color }) =>
    css`
      background-image: linear-gradient(
        -90deg,
        ${color ? `var(--${color})` : "var(--violete)"},
        transparent
      );
    `};
`;

export const CardContainer = styled.article`
  margin: 10px 5px;
  border-radius: 8px;
  padding: 10px;
  background-color: var(--lightPink);
  box-shadow: 1px 1px 10px 0px rgba(194, 194, 194, 1),
    -1px -1px 10px 0px rgba(194, 194, 194, 1);
`;

export const CardContainerTitle = styled.h4`
  margin-top: 20px;
  padding-left: 12px;
  color: var(--lightBlue);
`;
