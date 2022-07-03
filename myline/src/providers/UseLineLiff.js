
import React, { useEffect, useState } from 'react'
import { useDispatch, useSelector } from 'react-redux';
import liff from '@line/liff';

import {
    API_URL,
    setLineUserIdToken
} from '../utils/ApiClient'

const UseLineLiff = () => {

    const dispatch = useDispatch();
    const [lineToken, setLineToken] = useState({ idtoken: '', accessToken: '' })
    // useEffect(()=>{
    //     setLineUserIdToken(idtoken)
    // },[idtoken])

    useEffect(() => {

        liff.init({
            liffId: API_URL.LIFF_ID,
        }).then(async () => {
            console.log("isLoggedIn : ", liff.isLoggedIn());
            if (liff.isLoggedIn()) {
                const idtoken = liff.getIDToken()
                const accessToken = liff.getAccessToken()
                //alert(`idtoken : ${idtoken}`)
                setLineToken({ idtoken: idtoken, accessToken: accessToken })
                console.log("%cID TOKEN : ", "color: green", idtoken);
                liff.getProfile()
                    .then((profile) => {
                        console.log("%cPROFILE : ", "color: green", profile);
                    })
                    .catch((err) => {
                        console.log("%cERROR DETAIL", "color: red", err);
                    });
            } else {
                liff.login();
            }
        }).catch((err) => {
            console.log(
                "%cERROR HEADER",
                "color: red",
                err.code,
                err.message
            );
        });
        return () => { };
    }, [liff]);

    return lineToken

}

export default UseLineLiff