"use client";

import dynamic from "next/dynamic";
import "react-quill/dist/quill.snow.css";
import { RichWrapper } from "./RichText.sty";

const QuillEditor = dynamic(() => import("react-quill"), { ssr: false });

const quillModules = {
  toolbar: [
    [
      { header: [1, 2, 3, false] },
      { font: [] },
      { size: [] },
      "bold",
      "italic",
      { list: "ordered" },
      { list: "bullet" },
      { align: [] },
    ],
  ],
};

const quillFormats = [
  "header",
  "font",
  "size",
  "bold",
  "italic",
  "list",
  "bullet",
  "align",
];

const RichText = ({
  id,
  value,
  handleChangeRich,
}: {
  id: string | string[];
  value: string;
  handleChangeRich: (name: string | string[], value: string) => void;
}) => {
  return (
    <RichWrapper>
      <QuillEditor
        id={Array.isArray(id) ? id.join(".") : id}
        value={value}
        onChange={(val) => handleChangeRich(id, val)}
        modules={quillModules}
        formats={quillFormats}
      />
    </RichWrapper>
  );
};

export default RichText;
