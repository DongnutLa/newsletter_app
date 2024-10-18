import { cookies } from "next/headers";
import LoginContainer from "./container/LoginContainer";
import { redirect } from "next/navigation";

const LoginPage = async () => {
  const token = cookies().get("_auth_")?.value;
  if (token) {
    redirect("/admin");
  }

  return <LoginContainer />;
};

export default LoginPage;
