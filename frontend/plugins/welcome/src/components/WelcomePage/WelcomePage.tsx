import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import Button from '@material-ui/core/Button';
import Timer from '../Timer';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
 
const WelcomePage: FC<{}> = () => {
 const profile = { givenName: 'ระบบทุนการศึกษา' };
 
 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
      >      <Timer /></Header>
           <Content>
       <ContentHeader title="ระบบยื่นคำร้องขอรับทุนการศึกษา">
         <Link component={RouterLink} to="/create">
           <Button variant="contained" color="secondary">
             ยื่นคำร้อง
           </Button>
         </Link>
       </ContentHeader>
     </Content>
   </Page>
 );
};
 
export default WelcomePage;
