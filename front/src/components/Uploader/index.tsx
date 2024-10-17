import React, { FC, useCallback, useEffect, useState } from "react";
import { UploadFile, RcFile } from "antd/es/upload";
import Upload from "antd/es/upload/Upload";
import message from "antd/es/message";
import { useTranslations } from "next-intl";
import { useMutation } from "@tanstack/react-query";
import { uploadFile } from "@/lib/services/file";

interface UploaderProps {
  id: string;
  folder: string;
  fileName?: string;
  data: string[];
  maxFiles?: number;
  onUploadImage: (params: string[]) => void;
  onDeletedImage: (params: string[]) => void;
}

const Uploader: FC<UploaderProps> = ({
  id,
  folder,
  fileName,
  data,
  maxFiles = 1,
  onUploadImage,
  onDeletedImage,
}) => {
  const t = useTranslations();
  const [fileValue, setFileValue] = useState<string[]>(data);
  const [fileList, setFileList] = useState<UploadFile<Image>[]>([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (data?.length) {
      setFileValue(data);
      setFileList(
        data.map((i) => ({
          uid: i,
          name: i,
          size: 1000,
          status: "done",
          url: i,
        }))
      );
    }
  }, [data]);

  const { mutateAsync: UploadFile } = useMutation({
    mutationFn: uploadFile,
    onSuccess(data: string) {
      onUploadImage([...fileValue, data]);
      setFileValue([...fileValue, data]);
      setFileList([
        ...fileList,
        {
          uid: data,
          name: data,
          size: 1000,
          status: "done",
          url: data,
        },
      ]);
    },
  });

  const onRemove = useCallback(
    async (file: UploadFile<string>) => {
      let fileId = "";
      const fileIdx = fileValue.findIndex((f) => f === file?.url);
      if (fileIdx > -1) {
        fileId = fileValue[fileIdx];
      } else {
        fileId = file?.response ?? "";
      }

      if (!fileId) return false;

      const newFileValue = fileValue.filter((f) => f !== fileId);
      const newFileList = fileList.filter((f) => f.uid !== fileId);

      onDeletedImage(newFileValue);
      setFileValue(newFileValue);
      setFileList(newFileList);
    },
    [fileValue, fileList, onDeletedImage]
  );

  const beforeUpload = useCallback(
    async (file: RcFile): Promise<false> => {
      const isLt4M = file.size / 1024 / 1024 < 4;
      if (!isLt4M) {
        message.error("File must smaller than 2MB!");
        return false;
      }
      setLoading(true);

      const formData = new FormData();
      formData.append("file", file);
      formData.append("folder", folder);
      formData.append("fileName", fileName ?? "");

      await UploadFile(formData);
      setLoading(false);

      return false;
    },
    [UploadFile, folder, fileName, setLoading]
  );

  return (
    <Upload
      id={id}
      style={{
        width: "100%",
      }}
      accept="image/png, application/pdf"
      maxCount={maxFiles}
      listType="picture-card"
      fileList={fileList}
      beforeUpload={beforeUpload}
      onRemove={onRemove}
    >
      {fileList.length < maxFiles && <UploadButton loading={loading} t={t} />}
    </Upload>
  );
};

export default Uploader;

import { LoadingOutlined, PlusOutlined } from "@ant-design/icons";
import { Image } from "@/lib/models";
const UploadButton = ({
  loading,
  t,
}: {
  loading: boolean;
  t: (...args0: any) => string;
}) => (
  <div>
    {loading ? <LoadingOutlined /> : <PlusOutlined />}
    <div style={{ marginTop: 8 }}>{t("UploadImage")}</div>
  </div>
);
