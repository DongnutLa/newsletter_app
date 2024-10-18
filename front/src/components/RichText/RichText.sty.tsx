import styled from "styled-components";

export const RichWrapper = styled.div`
  background-color: white;
  border-radius: 6px;
  max-width: 350px;

  & .ql-toolbar {
    border-radius: 6px;
    box-shadow: 0px 2px 6px 0px #ccc;
  }

  & .ql-container {
    border-radius: 0 0 6px 6px;
    height: 200px;
  }
`;
