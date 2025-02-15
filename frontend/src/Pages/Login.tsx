import React, { Ref, RefObject, useEffect, useRef, useState } from "react"
import styles from './styles.module.scss'
import emailIcon from './../assets/emailIcon.svg'
import passwordIcon from './../assets/passwordIcon.svg'
import poster from './../assets/Parking.png'
import mobileLogo from './../assets/Website_logo.png'

function validateLogin(emailValue:RefObject<HTMLInputElement>,passwordValue:RefObject<HTMLInputElement>) {
  if(emailValue.current.value == null || passwordValue.current.value==null)
  {
    alert('please enter all fields')
  }
  if(!emailValue.current.value.includes('@',3) || passwordValue.current.value.length<=4){
    alert('either password is less than 5 charcaters or incorrect value entered for email')
  }
}
export function Login() {

  const emailValue=useRef<HTMLInputElement>(null);
  const passwordValue=useRef<HTMLInputElement>(null);
  const [width, setWidth] = useState<number>(window.innerWidth);

function handleWindowSizeChange() {
    setWidth(window.innerWidth);
}
useEffect(() => {
    window.addEventListener('resize', handleWindowSizeChange);
    return () => {
        window.removeEventListener('resize', handleWindowSizeChange);
    }
}, []);

const isMobile = width <= 720;
  return (
   
    <div className={styles.LoginContainer}>
       {
      isMobile&&
      <img  className={styles.mobile_Logo}src={mobileLogo}/>
    }
      <div className={styles.login_Input_Block}>
        <div className="email">
          <h2>Login to access your account</h2>
          <label>
            <img src={emailIcon} />
            <input type="input" placeholder="Your Email" ref={emailValue} required/>
          </label>
        </div>
        <div className="password">
          <label>
            <img src={passwordIcon} />
            <input type="password" placeholder="Your Password" ref={passwordValue}required />
          </label>
        </div>
        <button type="submit" className={styles.loginButton} onClick={()=>{validateLogin(emailValue,passwordValue)}}>Log in</button>
      </div>
      <div className={styles.login_Banner}>
        <img src={poster} />
      </div>
    </div>
  )
}