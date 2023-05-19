#include <termios.h>
#include "header.h"

char* USERS = "./data/users.txt";

void loginMenu(struct User* u)
{
    struct termios oflags, nflags;

    system("clear");
    printf("\n\n\n\t\t\t\t   Bank Management System\n\t\t\t\t\t User Login:");
    scanf("%s", u->name);

    // disabling echo
    tcgetattr(fileno(stdin), &oflags);
    nflags = oflags;
    nflags.c_lflag &= ~ECHO;
    nflags.c_lflag |= ECHONL;

    if (tcsetattr(fileno(stdin), TCSANOW, &nflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
    printf("\n\n\n\n\n\t\t\t\tEnter the password to login:");
    scanf("%s", u->password);

    // restore terminal
    if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
};


void registerMenu(struct User* u) {
    FILE* uf = fopen(USERS, "r+");

    printf("%d", sizeof * uf);

    struct termios oflags, nflags;

    system("clear");

    printf("\n\n\n\t\t\t\t   Bank Management System\n\t\t\t\t\t User Name:");
    scanf("%s", u->name);
    if (userExist(uf, u->name) == 1)
    {
        printf("✖ User with this name already exists\n\n");
        sleep(1);
            system("clear");
            
        initMenu(&u);
    }



    // disabling echo
    tcgetattr(fileno(stdin), &oflags);
    nflags = oflags;
    nflags.c_lflag &= ~ECHO;
    nflags.c_lflag |= ECHONL;

    if (tcsetattr(fileno(stdin), TCSANOW, &nflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
choosePass:
    printf("\n\n\n\n\n\t\t\t\tChoose a password:");
    scanf("%s", u->password);
    printf("\n\n\n\n\n\t\t\t\tConfirm your password:");

    char tPass[50];
    scanf("%s", tPass);
    if (strcmp(u->password, tPass) != 0)
    {
        printf("passwords doesn't match");
        goto choosePass;
    }
    else {
        printf("\nregistered\n");
        u->id = getLastUserIndex() + 1;
        saveUserToFile(uf, u);
        fclose(uf);
    }
    // restore terminal
    if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
    {
        perror("tcsetattr");
        return exit(1);
    }
};


const char* getPassword(struct User* u)
{
    FILE* fp = fopen(USERS, "r");
    struct User userChecker;

    if (fp == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    char line[100];
    while (fgets(line, sizeof(line), fp) != NULL)
    {
        char* token = strtok(line, " ");
        u->id = atoi(token);

        token = strtok(NULL, " ");
        if (strcmp(token, u->name) == 0)
        {
            fclose(fp);
            token = strtok(NULL, " ");
            token = strtok(token, "\n");
            return token;
        }
    }

    fclose(fp);
    return "no user found";
};

int getLastUserIndex()
{
    FILE* fp = fopen(USERS, "r");
    fseek(fp, 0, SEEK_END);
    long size = ftell(fp);
    int index;
    if (fp == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
    char line[100];

    if (size != 0)
    {
        fseek(fp,0,SEEK_SET);
        while (fgets(line, sizeof(line), fp) != NULL)
        {
            char* token = strtok(line, " ");
            index = atoi(token);
        }
        fclose(fp);
            return index;
    }
        fclose(fp);
        return 0;
};


void saveUserToFile(FILE* ptr, struct User* u)
{
    fprintf(ptr, "%d %s %s\n",
        u->id,
        &u->name,
        &u->password);
}

int userExist(FILE* ptr, char userName[50])
{
    char line[100];

    if (ptr != NULL)
    {

            while (fgets(line, sizeof(line), ptr) != NULL) {
                char* token = strtok(line, " ");
                token = strtok(NULL, " ");

                if (token != NULL && strcmp(token, userName) == 0) {
                    return 1;
                }
            }

    }
    return 0;

}

int getUserFromFile(FILE* ptr, struct User* u)
{
    if (ptr != NULL)
    {
        long size = ftell(ptr);

        if (0 != size) {
            return fscanf(ptr, "%d %s %s",
                &u->id,
                u->name,
                u->password) != EOF;
        }
    }
}