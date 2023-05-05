#include <termios.h>
#include "header.h"

char *USERS = "./data/users.txt";

void loginMenu(struct User *u)
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


void registerMenu(struct User *u){
    FILE * uf = fopen(USERS,"a+");

    struct termios oflags, nflags;

    system("clear");
    invalid:

    printf("\n\n\n\t\t\t\t   Bank Management System\n\t\t\t\t\t User Name:");
        scanf("%s", u->name);
            if (userExist(uf, u->name) == 1)
            {
                printf("✖ User with this name already exists\n\n");
                goto invalid;
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
        if (strcmp(u->password,tPass) != 0)
        {
            printf("passwords doesn't match");
            goto choosePass;
        } else {
            u->id = getLastUserIndex(USERS)+1;
            saveUserToFile(uf,u);
            fclose(uf);
        }
        // restore terminal
        if (tcsetattr(fileno(stdin), TCSANOW, &oflags) != 0)
        {
            perror("tcsetattr");
            return exit(1);
        }
};


const char *getPassword(struct User *u)
{
    FILE *fp = fopen("./data/users.txt", "r");
    struct User userChecker;

    if (fp == NULL)
    {
        printf("Error! opening file");
        exit(1);
    }
char line[100];
    while (fgets(line, sizeof(line), fp) != NULL)
    {
        char * token = strtok(line," ");
                    u->id = atoi(token);

        token = strtok(NULL," ");
        if (strcmp(token, u->name) == 0)
        {
            fclose(fp);
            token = strtok(NULL," ");
            token = strtok(token,"\n");
            return token;
        }
    }

    fclose(fp);
    return "no user found";
};

int getLastUserIndex(char* filename) {
    FILE* fp = fopen(filename, "r");
    if (fp == NULL) {
        printf("Failed to open file %s\n", filename);
        return -1;
    }

    char line[100];
    char* second_last_line = NULL;
    while (fgets(line, sizeof(line), fp) != NULL) {
        if (second_last_line != NULL) {
            free(second_last_line);  // free memory from previous iteration
        }
        second_last_line = strdup(line);  // make a copy of the current line
    }
    printf("SECOND: %s",second_last_line);
    fclose(fp);
    char  n = second_last_line[0];

    int index = atoi(&n);
    free(second_last_line);  // free memory from last iteration
    return index;
};

void saveUserToFile(FILE *ptr, struct User *u)
{
    fprintf(ptr, "%d %s %s\n",
	    u->id,
	    &u->name,
        &u->password);
}

int userExist(char userName[50])
{
    FILE *ptr = fopen(USERS,"r+");
    char line[100];

    while(fgets(line, sizeof(line),ptr) != NULL) {
        char* token = strtok(line, " ");
   token = strtok(NULL, " ");

        if (token != NULL && strcmp(token, userName) == 0) {
            return 1;
        }
    }
    fclose(ptr);
    return 0;

}

int getUserFromFile(FILE* ptr, struct User* u)
{
    return fscanf(ptr, "%d %s %s",
        &u->id,
        u->name,
        u->password) != EOF;
}