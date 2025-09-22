import { createContext, useContext, useEffect, useState } from "react"
import { useLocation, useNavigate } from "react-router-dom"


const Context  = createContext<{
    atk: string, 
    setAtkFunc: (newToken: string) => void,
    logout: () => void
}>({
    atk: "", 
    setAtkFunc: () => {},
    logout: () => {}
})
const AuthContextProvider = ({children}: {children: React.ReactNode}) => {
  const [atk, setAtk] = useState("")
  const path = useLocation().pathname
  const nav = useNavigate()




  useEffect(() => {
    const savedToken = localStorage.getItem("atk");
    if (savedToken) {
      setAtk(savedToken);
    }
    else if(!path.includes("/auth/") && !atk){
        nav("/auth/login")
        return
    } 
    else return
  }, [atk, path, nav]);

  const setAtkFunc = (newToken: string) => {
    if (newToken) {
        setAtk(newToken);
        localStorage.setItem("atk", newToken);
    } else {
        localStorage.removeItem("atk");
        setAtk("");
    }
  };

  const logout = () => {
    setAtkFunc("")
  }

  


  if(!path.includes("/auth/") && !atk){
    return null
  } 
  return (
    <Context.Provider value={{atk, setAtkFunc, logout}}>
        {children}
    </Context.Provider>
  )
}

export default  AuthContextProvider

export const useAuthContext = () => {
  return useContext(Context)
}