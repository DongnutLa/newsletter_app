import Home from "@/components/Home";

export default async function HomePage() {
  const newsletters = await getInitialData();

  return <Home newsletters={newsletters} />;
}

const getInitialData = async () => {
  // Get newsletters
  return [];
};
