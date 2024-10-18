import CreateNewsletterContainer from "../container/CreateNewsletterContainer";
import { cookies } from "next/headers";
import { listTopics } from "@/lib/services/topics";

const CreatePage = async () => {
  const topics = await getInitialData();
  return <CreateNewsletterContainer topics={topics} />;
};

export default CreatePage;

const getInitialData = async () => {
  const token = cookies().get("_auth_")?.value;

  return await listTopics(token);
};
