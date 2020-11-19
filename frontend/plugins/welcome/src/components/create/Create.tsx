import React, {useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import { Content, Header, Page, pageTheme, ContentHeader, } from '@backstage/core';

import {

  //Container,
  //Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  //TextField,
  //Avatar,
  Link,
  Button,
} from '@material-ui/core';
import Timer from '../Timer';
import { DefaultApi } from '../../api/apis'; // Api Gennerate From Command

import {EntUser} from '../../api/models/EntUser';
import {EntScholarshiptype} from '../../api/models/EntScholarshiptype';
import {EntScholarshipinformation} from '../../api/models/EntScholarshipinformation';
import {EntSemester} from '../../api/models/EntSemester';
import {EntScholarshipRequest} from '../../api/models/EntScholarshipRequest';
import { Alert } from '@material-ui/lab';


import { Link as RouterLink } from 'react-router-dom';
/*
const HeaderCustom = {
  minHeight: '50px',
};
*/
// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
    display: 'flex',
    justifyContent: 'center',
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(2),
  },
  formControl: {
    margin: theme.spacing(3),
    width: 320,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },

  margin: {
    margin: theme.spacing(3),
  },
  input: {
    display: 'none',
  },
}));

export default function CreateScholarshiprequest() {
  const classes = useStyles();
  const profile = { givenName: 'การยื่นคำร้องขอรับทุนการศึกษา' };
  const http = new DefaultApi();

  const [status, setStatus] = React.useState(false);
  const [alert, setAlert] = React.useState(true);

    const [ScholarshipRequest, setScholarshipRequest] = React.useState<EntScholarshipRequest[]>([]);
    const [User, setUser] =  React.useState<EntUser[]>([]);
    const [Scholarshiptype, setScholarshiptype] = React.useState<EntScholarshiptype[]>([]);
    const [Scholarshipinformation, setScholarshipinformation] = React.useState<EntScholarshipinformation[]>([]);
    const [Semester, setSemester] = React.useState<EntSemester[]>([]);
    
    const [Userid, setUserId] = React.useState(Number);
    const [Scholarshiptypeid, setScholarshiptypeId] = React.useState(Number);
    const [Scholarshipinformationid, setScholarshipinformationId] = React.useState(Number);
    const [Semesterid, setSemesterId] = React.useState(Number);

    useEffect(() => {
      const getUsers = async () => {
        const res = await http.listUser({ limit: 10, offset: 0 });
        setUser(res);
      };

      const getScholarshiptype = async () => {
        const res = await http.listScholarshiptype({ limit: 10, offset: 0 });
        setScholarshiptype(res);
      };

      const getScholarshipinformation = async () => {
        const res = await http.listScholarshipinformation({ limit: 10, offset: 0 });
        setScholarshipinformation(res);
      };

      const getSemester = async () => {
        const res = await http.listSemester({ limit: 10, offset: 0 });
        setSemester(res);
      };
      // Lifecycle Hooks
        getUsers();
        getScholarshiptype();
        getScholarshipinformation();
        getSemester();
    }, []);
    const getScholarshipRequest = async () => {
      const res = await http.listScholarshiprequest({ limit: 10, offset: 0 });
      setScholarshipRequest(res);
    }
    // set data to object Rent
    const UserhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setUserId(event.target.value as number);
    };
    const ScholarshiptypehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setScholarshiptypeId(event.target.value as number);
    };
    const ScholarshipinformationhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setScholarshipinformationId(event.target.value as number);
    };
    const SemesterhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setSemesterId(event.target.value as number);
    };
    const CreateScholarshiprequest = async () => {
      const Scholarshiprequest = {
        user : Userid,
        scholarshiptype : Scholarshiptypeid,
        scholarshipinformation : Scholarshipinformationid,
        semester : Semesterid,
      };
      console.log(Scholarshiprequest);
      const res: any = await http.createScholarshiprequest({ scholarshiprequest : Scholarshiprequest });
      setStatus(true);

      if (res.id != ''){
        setAlert(true);
      } else {
        setAlert(false);
      }
  
  };
  


 return (
   <Page theme={pageTheme.home}>
     <Header
       title={`${profile.givenName || 'to Backstage'}`}
       subtitle="Some quick intro and links."
     ><Timer /></Header>
     <Content>
       <ContentHeader title=" ">
         
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกข้อมูลการยื่นคำร้องขอรับทุนการศึกษา
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 ไม่สามารถบันทึกข้อมูลได้
               </Alert>
             )}
           </div>
         ) : null}
       </ContentHeader>

       <div className={classes.root}>

         
         <form noValidate autoComplete="off">

          <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>Email</InputLabel>
                <Select
                  name="Email"
                  value={Userid}
                  onChange={UserhandleChange}
                >
                  {User.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.email}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>


        </form>
       </div>
        
       <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ประเภททุนการศึกษา</InputLabel>
                <Select
                  name="room"
                  value={Scholarshiptypeid}
                  onChange={ScholarshiptypehandleChange}
                >
                  {Scholarshiptype.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.typeName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
             </form>
          </div>
         
          <div className={classes.root}>
            <form noValidate autoComplete="off">
 
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ทุนการศึกษา</InputLabel>
                <Select
                  name="insurance"
                  value={Scholarshipinformationid}
                  onChange={ScholarshipinformationhandleChange}
                >
                  {Scholarshipinformation.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.scholarshipName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>

           </form>
       </div>





        <div className={classes.root}>
         <form noValidate autoComplete="off">

         <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>ภาคการศึกษา</InputLabel>
                <Select
                 name="user"
                  value={Semesterid}
                  onChange={SemesterhandleChange}
                >
                  {Semester.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.semester}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
        
         
         

        


        </form>
        
       </div>


       <div className={classes.root}>
         <form noValidate autoComplete="off">
           
         
           <div className={classes.margin}>
             <Button
               onClick={() => {
                CreateScholarshiprequest();
               }}
               variant="contained"
               color="primary"
             >
               ยืนยันข้อมูลคำร้องขอรับทุนการศึกษา
             </Button>
              
         <Link component={RouterLink} to="/">
           <Button variant="contained" color="secondary">
           กลับสู่หน้าหลัก
           </Button>
         </Link>

           </div>
         </form>
       </div>
  







     </Content>
   </Page>
 );
}
