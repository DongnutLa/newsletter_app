import HomeContainer from "@/container/HomeContainer";
import { listTopics } from "@/lib/services/topics";

export default async function HomePage() {
  const topics = await getInitialData();

  return <HomeContainer topics={topics} />;
}

const getInitialData = async () => {
  const topics = await listTopics();

  return topics;
};
