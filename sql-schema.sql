USE [tda]
GO

/****** Object:  Table [dbo].[tblContacts]    Script Date: 29/06/2020 4:58:35 PM ******/
DROP TABLE [dbo].[tblContacts]
GO

/****** Object:  Table [dbo].[tblContacts]    Script Date: 29/06/2020 4:58:35 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[tblContacts](
	[contactId] [varchar](255) NOT NULL,
	[fullName] [varchar](255) NOT NULL,
	[email] [varchar](255) NULL,
 CONSTRAINT [PK_contacts] PRIMARY KEY CLUSTERED 
(
	[contactId] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[tblPhoneNumbers] DROP CONSTRAINT [fk_contacts_phone_numbers]
GO

/****** Object:  Table [dbo].[tblPhoneNumbers]    Script Date: 29/06/2020 4:58:52 PM ******/
DROP TABLE [dbo].[tblPhoneNumbers]
GO

/****** Object:  Table [dbo].[tblPhoneNumbers]    Script Date: 29/06/2020 4:58:52 PM ******/
SET ANSI_NULLS ON
GO

SET QUOTED_IDENTIFIER ON
GO

CREATE TABLE [dbo].[tblPhoneNumbers](
	[phoneId] [varchar](255) NOT NULL,
	[contactId] [varchar](255) NOT NULL,
	[phoneNumber] [varchar](20) NOT NULL,
 CONSTRAINT [PK_phone_numbers] PRIMARY KEY CLUSTERED 
(
	[phoneId] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, IGNORE_DUP_KEY = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
) ON [PRIMARY]
GO

ALTER TABLE [dbo].[tblPhoneNumbers]  WITH CHECK ADD  CONSTRAINT [fk_contacts_phone_numbers] FOREIGN KEY([contactId])
REFERENCES [dbo].[tblContacts] ([contactId])
ON UPDATE CASCADE
ON DELETE CASCADE
GO