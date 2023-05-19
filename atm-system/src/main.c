#include "header.h"

void mainMenu(struct User* u)
{
    int op= NULL;

    system("clear");
    printf("userId: %d", u->id);
    printf("\n\n\t\t======= ATM =======\n\n");
    printf("\n\t\t-->> Feel free to choose one of the options below <<--\n");
    printf("\n\t\t[1]- Create a new account\n");
    printf("\n\t\t[2]- Update account information\n");
    printf("\n\t\t[3]- Check accounts\n");
    printf("\n\t\t[4]- Check list of owned account\n");
    printf("\n\t\t[5]- Make Transaction\n");
    printf("\n\t\t[6]- Remove existing account\n");
    printf("\n\t\t[7]- Transfer ownership\n");
    printf("\n\t\t[8]- Exit\n");

    scanf("%d", &op);
    getchar();



    switch (op)
    {
    case 1:
        createNewAcc(u);
        break;
    case 2:
        updateAccount(u);
        break;
    case 3:
        checkAccountDetails(u);
        break;
    case 4:
        checkAllAccounts(u);
        break;
    case 5:
        makeTransaction(u);
        break;
    case 6:
        removeAccount(u);
        break;
    case 7:
        transferOwnership(u);
        break;
    case 8:
        exit(1);
        break;
    default:
        invalidOperation(mainMenu,u);
    }
};

void initMenu(struct User* u)
{

    int r = 0;
    static int opt = NULL;
    system("clear");
    printf("\n\n\t\t======= ATM =======\n");
    printf("\n\t\t-->> Feel free to login / register :\n");
    printf("\n\t\t[1]- login\n");
    printf("\n\t\t[2]- register\n");
    printf("\n\t\t[3]- exit\n");
    scanf("%d", &opt);
    getchar();
    switch (opt)
    {
    case 1:
    wrongPassword:
        loginMenu(u);
        char * name = u->name;
        char* b = getPassword(u);

        if (strcmp(u->password, b) != 0)
        {
            printf("\nWrong password!! or User Name\n");
            sleep(1);
           goto wrongPassword;
        }
        r = 1;
        break;
    case 2:
        registerMenu(u);
        r = 1;
        break;
    case 3:
        exit(1);
        break;
    default:
       invalidOperation(initMenu,u);
    }
    
};

int main()
{
    struct User u ; //= {
    //     id:2,
    //     name:"Billy",
    //     password:"azerty"
    // };

    initMenu(&u);
    mainMenu(&u);

    return 0;
}
