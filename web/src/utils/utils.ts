import {User} from "@/api/user";
import {Md5} from "ts-md5";

export const getAvatarUrl = (user:User): string =>{
  console.log(user)
  if (user.avatar) {
    return user.avatar
  }
  return `https://cravatar.cn/avatar/${Md5.hashStr(user.email)}.png`
}

export const loginRedirectUrl = () => {
  return `https://auth.phy.pku.edu.cn/oidc/authorize/?response_type=code&scope=openid profile email address pku&client_id=16302204390022&redirect_uri=${import.meta.env.VITE_HOST_URL}/login`
}
